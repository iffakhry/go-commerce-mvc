package requests

import "github.com/iffakhry/go-commerce-mvc/entities"

type ProductRequest struct {
	UserID      uint    `json:"user_id" form:"user_id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Stock       uint    `json:"stock" form:"stock"`
	Description string  `json:"description" form:"description"`
}

func ProductRequestToEntity(dataRequest ProductRequest) entities.Product {
	return entities.Product{
		UserID:      dataRequest.UserID,
		Name:        dataRequest.Name,
		Price:       dataRequest.Price,
		Stock:       dataRequest.Stock,
		Description: dataRequest.Description,
	}
}
