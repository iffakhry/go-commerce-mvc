package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/iffakhry/go-commerce-mvc/controller/requests"
	"github.com/iffakhry/go-commerce-mvc/controller/responses"
	"github.com/iffakhry/go-commerce-mvc/entity"
	"github.com/iffakhry/go-commerce-mvc/pkg"
	"github.com/iffakhry/go-commerce-mvc/pkg/middlewares"
	"github.com/iffakhry/go-commerce-mvc/repository"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productRepo repository.ProductRepository
}

func NewProductController(repo repository.ProductRepository) *ProductController {
	return &ProductController{
		productRepo: repo,
	}
}

func (cp *ProductController) Create(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, pkg.FailedResponse("unauthorized"))
	}

	dataInput := requests.ProductRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error bind data"))
	}
	// mapping dari request ke core
	entityData := entity.Product{
		UserID:      uint(idToken), // ambil user id dari token yang login
		Name:        dataInput.Name,
		Price:       dataInput.Price,
		Stock:       dataInput.Stock,
		Description: dataInput.Description,
	}
	err := cp.productRepo.Insert(entityData)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, pkg.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, pkg.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusCreated, pkg.SuccessResponse("success insert data"))
}

func (cp *ProductController) GetAll(c echo.Context) error {
	// memanggil func di repositories
	results, err := cp.productRepo.SelectAll()

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
	var productsResponse = responses.ProductEntityToResponseList(results)

	// response ketika berhasil
	return c.JSON(http.StatusOK, pkg.SuccessWithDataResponse("success read data", productsResponse))
}

func (cp *ProductController) GetById(c echo.Context) error {
	idProduct := c.Param("id")
	idProductConv, errConv := strconv.Atoi(idProduct)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error read param id"))
	}

	// memanggil func di repositories
	result, err := cp.productRepo.SelectById(idProductConv)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.FailedResponse("error read data"))
	}

	var usersResponse = responses.ProductEntityToResponse(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, pkg.SuccessWithDataResponse("success read data", usersResponse))
}

func (cp *ProductController) Update(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, pkg.FailedResponse("unauthorized"))
	}

	idProduct := c.Param("id")
	idProductConv, errConv := strconv.Atoi(idProduct)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error read param id"))
	}
	// var idConv int
	// if id != "" {
	// 	var errConv error
	// 	idConv, errConv = strconv.Atoi(id)
	// 	if errConv != nil {
	// 		return c.JSON(http.StatusBadRequest, helper.FailedResponse("param id is required and must be number"))
	// 	}
	// }

	dataReq := requests.ProductRequest{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error bind data, please check your request body"))
	}
	dataEntity := requests.ProductRequestToEntity(dataReq)
	data, err := cp.productRepo.Update(idProductConv, dataEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, pkg.SuccessWithDataResponse("success update data", responses.ProductEntityToResponse(data)))
}

func (cp *ProductController) Delete(c echo.Context) error {

	idToken := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusUnauthorized, pkg.FailedResponse("unauthorized"))
	}

	idProduct := c.Param("id")
	idProductConv, errConv := strconv.Atoi(idProduct)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error read param id"))
	}

	row, err := cp.productRepo.Delete(idProductConv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, pkg.FailedResponse(err.Error()))
	}

	if row <= 0 {
		return c.JSON(http.StatusBadRequest, pkg.FailedResponse("error delete data"))
	}
	return c.JSON(http.StatusOK, pkg.SuccessResponse("success delete data"))
}
