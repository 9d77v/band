package main

import (
	"database/sql"
	"log"
	"strings"

	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
	"gorm.io/gorm"
)

// --- 枚举类型 ---

// enumTypeDef 枚举类型定义
type enumTypeDef struct {
	Name string
	SQL  string
}

func defaultEnumTypes() []enumTypeDef {
	return []enumTypeDef{
		{Name: "gender_type", SQL: "CREATE TYPE gender_type AS ENUM ('undefined', 'male', 'female')"},
		{Name: "order_type", SQL: "CREATE TYPE order_type AS ENUM ('undefined', 'coins', 'membership', 'goods')"},
		{Name: "order_status", SQL: "CREATE TYPE order_status AS ENUM ('undefined', 'pending', 'cancelled', 'completed', 'refunded')"},
		{Name: "payment_method", SQL: "CREATE TYPE payment_method AS ENUM ('undefined', 'wechat', 'alipay', 'apple_iap')"},
		{Name: "payment_platform", SQL: "CREATE TYPE payment_platform AS ENUM('undefined','web','wap','app','ios')"},
		{Name: "payment_status", SQL: "CREATE TYPE payment_status AS ENUM ('undefined', 'unpaid', 'in_progress', 'paid', 'refunded', 'failed')"},
		{Name: "product_type", SQL: "CREATE TYPE product_type AS ENUM ('physical', 'consumable', 'non_consumable', 'auto_renewable_subscription', 'non_renewing_subscription')"},
	}
}

// migrateEnumTypes 创建数据库枚举类型（不存在时创建）
func migrateEnumTypes(db *postgres.PgDB) {
	for _, et := range defaultEnumTypes() {
		exists, err := checkTypeExists(db.GetDB(), et.Name)
		if err != nil {
			log.Fatalf("检查枚举类型 %s 失败: %s", et.Name, err)
		}
		if exists {
			continue
		}
		err = db.Exec(et.SQL).Error()
		if err != nil && !strings.Contains(err.Error(), "SQLSTATE 42710") {
			log.Fatalf("创建枚举类型 %s 失败: %s", et.Name, err)
		}
	}
}

// checkTypeExists 检查数据库中是否存在指定的枚举类型
func checkTypeExists(db *gorm.DB, typeName string) (bool, error) {
	var exists bool
	err := db.Raw(`SELECT EXISTS (
		SELECT 1
		FROM pg_type t
		LEFT JOIN pg_namespace n ON n.oid = t.typnamespace
		WHERE t.typname = ?
		  AND t.typtype = 'e'
	)`, typeName).Scan(&exists).Error
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
