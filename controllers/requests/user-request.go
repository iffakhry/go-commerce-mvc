package requests

import "github.com/iffakhry/go-commerce-mvc/entities"

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Role     string `json:"role" form:"role"`
}

func UserRequestToEntity(dataRequest UserRequest) entities.User {
	return entities.User{
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
		Name:     dataRequest.Name,
		Role:     dataRequest.Role,
	}
}
