package models

import (
	"github.com/iffakhry/go-commerce-mvc/entities"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID      uint
	Name        string
	Price       float64
	Stock       uint
	Description string
	User        User
}

// mapping dari Entity ke model
func ProductEntityToModel(dataEntity entities.Product) Product {
	return Product{
		UserID:      dataEntity.UserID,
		Name:        dataEntity.Name,
		Price:       dataEntity.Price,
		Stock:       dataEntity.Stock,
		Description: dataEntity.Description,
		User:        UserEntityToModel(dataEntity.User),
	}
}

func ProductModelToEntity(dataModel Product) entities.Product {
	return entities.Product{
		Id:          dataModel.ID,
		UserID:      dataModel.UserID,
		Name:        dataModel.Name,
		Price:       dataModel.Price,
		Stock:       dataModel.Stock,
		Description: dataModel.Description,
		User:        UserModelToEntity(dataModel.User),
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}

func ProductModelToEntityList(dataModel []Product) []entities.Product {
	var coreList []entities.Product
	for _, v := range dataModel {
		coreList = append(coreList, ProductModelToEntity(v))
	}
	return coreList
}
