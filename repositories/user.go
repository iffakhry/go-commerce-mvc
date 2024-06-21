package repositories

import (
	"errors"
	"fmt"

	"github.com/iffakhry/go-commerce-mvc/entities"
	"github.com/iffakhry/go-commerce-mvc/models"
	"github.com/iffakhry/go-commerce-mvc/pkg"
	"github.com/iffakhry/go-commerce-mvc/pkg/middlewares"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Insert(input entities.User) error {
	// mapping dari struct entities core ke gorm model
	// userInputGorm := User{
	// 	Name:     input.Name,
	// 	Phone:    input.Phone,
	// 	Email:    input.Email,
	// 	Password: input.Password,
	// }
	hashedPassword, errHash := pkg.HashPassword(input.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	userInputGorm := models.UserEntityToModel(input)
	userInputGorm.Password = hashedPassword

	tx := repo.db.Create(&userInputGorm) // insert into users set name = .....
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

func (repo *UserRepository) SelectAll() ([]entities.User, error) {
	var usersData []models.User
	tx := repo.db.Find(&usersData) // select * from users
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(usersData)
	// mapping dari struct gorm model ke struct entities core
	var usersCoreAll []entities.User = models.UserModelToEntityList(usersData)
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
	return usersCoreAll, nil
}

func (repo *UserRepository) SelectById(id int) (entities.User, error) {
	var usersData models.User
	tx := repo.db.First(&usersData, id) // select * from users where id = id
	if tx.Error != nil {
		return entities.User{}, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var usersCore = models.UserModelToEntity(usersData)
	fmt.Println(usersCore)

	return usersCore, nil
}

func (repo *UserRepository) Update(id int, input entities.User) (data entities.User, err error) {
	if input.Password != "" {
		hashedPassword, errHash := pkg.HashPassword(input.Password)
		if errHash != nil {
			return entities.User{}, errors.New("error hash password")
		}
		input.Password = hashedPassword
	}

	tx := repo.db.Model(&models.User{}).Where("id = ?", id).Updates(models.UserEntityToModel(input))
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, errors.New("failed update data, row affected = 0")
	}
	var usersData models.User
	resultFind := repo.db.Find(&usersData, id)
	if resultFind.Error != nil {
		return entities.User{}, resultFind.Error
	}
	data = models.UserModelToEntity(usersData)
	return data, nil
}

func (repo *UserRepository) Delete(id int) (row int, err error) {
	result := repo.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed delete data. Data is not exits row = 0")
	}
	return int(result.RowsAffected), nil
}

func (repo *UserRepository) Login(email string, password string) (entities.User, string, error) {
	var userGorm models.User
	tx := repo.db.Where("email = ?", email).First(&userGorm) // select * from users limit 1
	if tx.Error != nil {
		return entities.User{}, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return entities.User{}, "", errors.New("login failed, email dan password salah")
	}

	checkPassword := pkg.CheckPasswordHash(password, userGorm.Password)
	if !checkPassword {
		return entities.User{}, "", errors.New("login failed, password salah")
	}

	token, errToken := middlewares.CreateToken(int(userGorm.ID))
	if errToken != nil {
		return entities.User{}, "", errToken
	}

	dataCore := models.UserModelToEntity(userGorm)
	return dataCore, token, nil
}
