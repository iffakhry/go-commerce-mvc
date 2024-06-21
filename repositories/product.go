package repositories

import (
	"errors"
	"fmt"

	"github.com/iffakhry/go-commerce-mvc/entities"
	"github.com/iffakhry/go-commerce-mvc/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) Insert(input entities.Product) error {
	// mapping dari struct entities core ke gorm model
	// userInputGorm := User{
	// 	Name:     input.Name,
	// 	Phone:    input.Phone,
	// 	Email:    input.Email,
	// 	Password: input.Password,
	// }
	dataInputGorm := models.ProductEntityToModel(input)

	tx := repo.db.Create(&dataInputGorm) // insert into product set ...
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

func (repo *ProductRepository) SelectAll() ([]entities.Product, error) {
	var datas []models.Product
	tx := repo.db.Preload("User").Find(&datas) // select * from users
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(datas)
	// mapping dari struct gorm model ke struct entities core
	var datasEntityAll []entities.Product = models.ProductModelToEntityList(datas)
	// for _, value := range usersData {
	// 	var userCore = user.Core{
	// 		Id:        value.ID,
	// 		Name:      value.Name,
	// 		Phone:     value.Phone,
	// 		Email:     value.Email,
	// 		Password:  value.Password,
	// 		CreatedAt: value.CreatedAt,
	// 		UpdatedAt: value.UpdatedAt,
	// 	}
	// 	usersCoreAll = append(usersCoreAll, userCore)
	// }
	return datasEntityAll, nil
}

func (repo *ProductRepository) SelectById(id int) (entities.Product, error) {
	var datas models.Product
	tx := repo.db.Preload("User").First(&datas, id) // select * from users where id = id
	if tx.Error != nil {
		return entities.Product{}, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var core = models.ProductModelToEntity(datas)
	fmt.Println(core)

	return core, nil
}

func (repo *ProductRepository) Update(id int, input entities.Product) (data entities.Product, err error) {

	tx := repo.db.Model(&models.Product{}).Where("id = ?", id).Updates(models.ProductEntityToModel(input))
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, errors.New("failed update data, row affected = 0")
	}
	var dataModel models.Product
	resultFind := repo.db.Preload("User").Find(&dataModel, id)
	if resultFind.Error != nil {
		return entities.Product{}, resultFind.Error
	}
	data = models.ProductModelToEntity(dataModel)
	return data, nil
}

func (repo *ProductRepository) Delete(id int) (row int, err error) {
	result := repo.db.Delete(&models.Product{}, id)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed delete data. Data is not exits row = 0")
	}
	return int(result.RowsAffected), nil
}
