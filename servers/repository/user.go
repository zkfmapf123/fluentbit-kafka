package repository

import (
	"fmt"

	"github.com/zkfmapf123/fpg/config"
	"github.com/zkfmapf123/fpg/models"
)

type UserRepository struct {
	db *config.DBConn
}

func NewUserRepository(db *config.DBConn) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *models.User) error {

	res := ur.db.DB.Create(user)

	fmt.Println(res)

	return nil

}
