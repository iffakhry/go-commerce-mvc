package responses

import (
	"time"

	"github.com/iffakhry/go-commerce-mvc/entity"
)

type UserResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func UserEntityToResponse(dataCore entity.User) UserResponse {
	return UserResponse{
		Id:        dataCore.Id,
		Name:      dataCore.Name,
		Email:     dataCore.Email,
		Role:      dataCore.Role,
		CreatedAt: dataCore.CreatedAt,
	}
}

func UserEntityToResponseList(dataCore []entity.User) []UserResponse {
	var result []UserResponse
	for _, v := range dataCore {
		result = append(result, UserEntityToResponse(v))
	}
	return result
}
