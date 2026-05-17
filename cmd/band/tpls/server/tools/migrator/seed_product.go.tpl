package main

import (
	"encoding/json"
	"fmt"
	"log"

	productEnum "hhy-services/apps/product/domain/enum"
	productDo "hhy-services/apps/product/infrastructure/persistence/postgres/do"

	"github.com/9d77v/band/pkg/app"
	"github.com/9d77v/band/pkg/stores/orm/base"
	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
	"gorm.io/gorm"
)

var uniqueID = app.NewUniqueID(app.Conf{})

// --- 种子数据: 商品 ---

// coinPackage 金币套餐定义
type coinPackage struct {
	PurchasedAmount int64
	GiftAmount      int64
	Price           int64
}

func defaultCoinPackages() []coinPackage {
	return []coinPackage{
		{PurchasedAmount: 600, GiftAmount: 10, Price: 600},
		{PurchasedAmount: 1000, GiftAmount: 20, Price: 1000},
		{PurchasedAmount: 3000, GiftAmount: 100, Price: 3000},
		{PurchasedAmount: 5000, GiftAmount: 300, Price: 5000},
		{PurchasedAmount: 10000, GiftAmount: 1000, Price: 10000},
		{PurchasedAmount: 20000, GiftAmount: 2700, Price: 20000},
	}
}

// seedProducts 初始化商品种子数据
func seedProducts(db *postgres.PgDB) {
	var productCount int64
	err := db.Model(&productDo.Product{}).Where("code = ?", "COIN_RECHARGE").Count(&productCount).Error()
	if err != nil {
		log.Fatalf("查询商品数量失败: %s", err)
	}
	if productCount > 0 {
		return
	}

	err = db.GetDB().Transaction(func(tx *gorm.DB) error {
		productID := uniqueID.GetID()
		desc := "金币充值套餐"
		product := &productDo.Product{
			Model:       base.Model{ID: productID},
			Name:        "金币套餐",
			Code:        "COIN_RECHARGE",
			ProductType: productEnum.ProductTypeConsumable,
			Description: &desc,
			IsActive:    true,
		}
		if err := tx.Create(product).Error; err != nil {
			return fmt.Errorf("创建商品失败: %w", err)
		}

		for i, p := range defaultCoinPackages() {
			attrs := map[string]int64{
				"purchased_coins": p.PurchasedAmount,
				"gift_coins":      p.GiftAmount,
			}
			attrsJSON, _ := json.Marshal(attrs)
			sku := &productDo.Sku{
				Model:      base.Model{ID: uniqueID.GetID()},
				ProductID:  productID,
				Name:       coinPackageName(p.PurchasedAmount, p.GiftAmount),
				Price:      int(p.Price),
				Currency:   "CNY",
				SortOrder:  i + 1,
				IsActive:   true,
				Attributes: attrsJSON,
			}
			if err := tx.Create(sku).Error; err != nil {
				return fmt.Errorf("创建 SKU 失败: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("商品种子数据创建失败: %s", err)
	}
	log.Println("商品和 SKU 种子数据创建成功")
}

// coinPackageName 生成金币套餐名称
func coinPackageName(purchasedCoins, giftCoins int64) string {
	if giftCoins > 0 {
		return fmt.Sprintf("%d金币(送%d)", purchasedCoins, giftCoins)
	}
	return fmt.Sprintf("%d金币", purchasedCoins)
}
