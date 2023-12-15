package audit

import (
	"github.com/9d77v/band/pkg/jwt"
	"github.com/9d77v/band/pkg/stores/orm/base"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditEntity struct {
	base.Entity
	CreatedBy uuid.UUID `json:"createdBy" gorm:"comment:创建人"` //创建人id
	UpdatedBy uuid.UUID `json:"updatedBy" gorm:"comment:更新人"` //更新人id
}

func (u *AuditEntity) BeforeCreate(tx *gorm.DB) (err error) {
	claims := jwt.NewContextKey("claims").Get(tx.Statement.Context)
	if claims != nil {
		tx.Statement.SetColumn("CreatedBy", claims.ID)
		tx.Statement.SetColumn("UpdatedBy", claims.ID)
	}
	return
}

func (u *AuditEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	claims := jwt.NewContextKey("claims").Get(tx.Statement.Context)
	if claims != nil {
		tx.Statement.SetColumn("UpdatedBy", claims.ID)
	}
	return
}

func (u *AuditEntity) BeforeDelete(tx *gorm.DB) (err error) {
	claims := jwt.NewContextKey("claims").Get(tx.Statement.Context)
	if claims != nil {
		tx.Statement.SetColumn("UpdatedBy", claims.ID)
	}
	return
}

func (u *AuditEntity) BeforeSave(tx *gorm.DB) (err error) {
	claims := jwt.NewContextKey("claims").Get(tx.Statement.Context)
	if claims != nil {
		tx.Statement.SetColumn("UpdatedBy", claims.ID)
		if u == nil {
			tx.Statement.SetColumn("CreatedBy", claims.ID)
		}
	}
	return
}
