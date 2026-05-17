package main

import (
	"log"

	aigcDo "hhy-services/apps/aigc/infrastructure/persistence/do"
	fileDo "hhy-services/apps/file/infrastructure/persistence/do"
	orderDo "hhy-services/apps/order/infrastructure/persistence/do"
	paymentDo "hhy-services/apps/payment/infrastructure/persistence/do"
	productDo "hhy-services/apps/product/infrastructure/persistence/postgres/do"
	speechDo "hhy-services/apps/speech/infrastructure/persistence/do"
	userDo "hhy-services/apps/user/infrastructure/persistence/do"

	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
)

// migrateSchemas 自动迁移数据库表结构
func migrateSchemas(db *postgres.PgDB) {
	err := db.GetDB().AutoMigrate(
		&userDo.User{},
		&userDo.UserAuth{},
		&userDo.UserCoin{},
		&userDo.UserCoinRecord{},
		&userDo.UserFeedback{},
		&fileDo.File{},
		&fileDo.FileAccount{},
		&fileDo.UserFile{},
		&aigcDo.AigcApiLog{},
		&aigcDo.AigcUserActivity{},
		&aigcDo.AigcPrompt{},
		&aigcDo.AigcTask{},
		&aigcDo.AigcTaskInvocationRecord{},
		&orderDo.Order{},
		&orderDo.OrderDetail{},
		&paymentDo.Payment{},
		&paymentDo.PaymentCallback{},
		&paymentDo.Refund{},
		&paymentDo.RefundCallback{},
		&productDo.Product{},
		&productDo.Sku{},
		&speechDo.ConversationScript{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}
