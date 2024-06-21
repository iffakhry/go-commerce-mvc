package models

import (
	"github.com/iffakhry/go-commerce-mvc/entity"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID        string `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Email    string `gorm:"unique"`
	Password string
	Name     string
	Role     string
}

// mapping dari Entity ke model
func UserEntityToModel(dataEntity entity.User) User {
	return User{
		Email:    dataEntity.Email,
		Password: dataEntity.Password,
		Name:     dataEntity.Name,
		Role:     dataEntity.Role,
	}
}

func UserModelToEntity(dataModel User) entity.User {
	return entity.User{
		Id:        dataModel.ID,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Name:      dataModel.Name,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func UserModelToEntityList(dataModel []User) []entity.User {
	var coreList []entity.User
	for _, v := range dataModel {
		coreList = append(coreList, UserModelToEntity(v))
	}
	return coreList
}
