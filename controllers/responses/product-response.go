package responses

import (
	"time"

	"github.com/iffakhry/go-commerce-mvc/entities"
)

type ProductResponse struct {
	Id          uint         `json:"id"`
	UserID      uint         `json:"user_id"`
	Name        string       `json:"name"`
	Price       float64      `json:"price"`
	Stock       uint         `json:"stock"`
	Description string       `json:"description"`
	User        UserResponse `json:"user,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func ProductEntityToResponse(dataCore entities.Product) ProductResponse {
	return ProductResponse{
		Id:          dataCore.Id,
		UserID:      dataCore.UserID,
		Name:        dataCore.Name,
		Price:       dataCore.Price,
		Stock:       dataCore.Stock,
		Description: dataCore.Description,
		User:        UserEntityToResponse(dataCore.User),
		CreatedAt:   dataCore.CreatedAt,
		UpdatedAt:   dataCore.UpdatedAt,
	}
}

func ProductEntityToResponseList(dataCore []entities.Product) []ProductResponse {
	var result []ProductResponse
	for _, v := range dataCore {
		result = append(result, ProductEntityToResponse(v))
	}
	return result
}
