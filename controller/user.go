package controller

import (
	"net/http"
	"strings"

	"github.com/iffakhry/go-commerce-mvc/controller/requests"
	"github.com/iffakhry/go-commerce-mvc/controller/responses"
	"github.com/iffakhry/go-commerce-mvc/entity"
	"github.com/iffakhry/go-commerce-mvc/pkg"
	"github.com/iffakhry/go-commerce-mvc/pkg/middlewares"
	"github.com/iffakhry/go-commerce-mvc/repository"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userRepo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{
		userRepo: repo,
	}
}

func (cuser *UserController) CreateUser(c echo.Context) error {
	userInput := requests.UserRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error bind data"))
	}
	// mapping dari request ke core
	userCore := entity.User{
		Email:    userInput.Email,
		Password: userInput.Password,
		Name:     userInput.Name,
		Role:     userInput.Role,
	}
	err := cuser.userRepo.Insert(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, pkg.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, pkg.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusCreated, pkg.SuccessResponse("success insert data"))
}

func (cuser *UserController) Login(c echo.Context) error {
	loginInput := requests.AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error bind data"))
	}

	dataLogin, token, err := cuser.userRepo.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, pkg.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, pkg.FailedResponse("error login, intrnal server error"))
		}
	}

	return c.JSON(http.StatusOK, pkg.SuccessWithDataResponse("login success", map[string]any{
		"token": token,
		"email": dataLogin.Email,
		"id":    dataLogin.Id,
	}))
}

func (cuser *UserController) GetAllUser(c echo.Context) error {
	// memanggil func di repositories
	results, err := cuser.userRepo.SelectAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error read data"))
	}

	// var usersResponse []responses.UserResponse
	// for _, value := range results {
	// 	usersResponse = append(usersResponse, responses.UserResponse{
	// 		Id:        value.Id,
	// 		Name:      value.Name,
	// 		Email:     value.Email,
	// 		CreatedAt: value.CreatedAt,
	// 	})
	// }
	var usersResponse = responses.UserEntityToResponseList(results)

	// response ketika berhasil
	return c.JSON(http.StatusOK, pkg.SuccessWithDataResponse("success read data", usersResponse))
}

func (cuser *UserController) GetProfile(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)

	// memanggil func di repositories
	result, err := cuser.userRepo.SelectById(idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.FailedResponse("error read data"))
	}

	var usersResponse = responses.UserEntityToResponse(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, pkg.SuccessWithDataResponse("success read profile", usersResponse))
}

func (cuser *UserController) Update(c echo.Context) error {
	// id := c.Param("id")
	// var idConv int
	// if id != "" {
	// 	var errConv error
	// 	idConv, errConv = strconv.Atoi(id)
	// 	if errConv != nil {
	// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("param id is required and must be number"))
	// 	}
	// }

	idToken := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, pkg.FailedResponse("unauthorized"))
	}

	userReq := requests.UserRequest{}
	errBind := c.Bind(&userReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error bind data, please check your request body"))
	}
	userCore := requests.UserRequestToEntity(userReq)
	data, err := cuser.userRepo.Update(idToken, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.SuccessWithDataResponse("success update data", responses.UserEntityToResponse(data)))
}

func (cuser *UserController) Delete(c echo.Context) error {

	idToken := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, pkg.FailedResponse("unauthorized"))
	}
	row, err := cuser.userRepo.Delete(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.FailedResponse(err.Error()))
	}

	if row <= 0 {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error delete data"))
	}
	return c.JSON(http.StatusOK, pkg.SuccessResponse("success deactivate account"))
}
