package models

import (
	"github.com/iffakhry/go-commerce-mvc/entities"
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
func UserEntityToModel(dataEntity entities.User) User {
	return User{
		Email:    dataEntity.Email,
		Password: dataEntity.Password,
		Name:     dataEntity.Name,
		Role:     dataEntity.Role,
	}
}

func UserModelToEntity(dataModel User) entities.User {
	return entities.User{
		Id:        dataModel.ID,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Name:      dataModel.Name,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func UserModelToEntityList(dataModel []User) []entities.User {
	var coreList []entities.User
	for _, v := range dataModel {
		coreList = append(coreList, UserModelToEntity(v))
	}
	return coreList
}
