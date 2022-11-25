package user

import (
	"database/sql"
	"fmt"

	"github.com/forsunforson/profolio/internal/model"
)

type Repository interface {
	GetUserByID(id int64) (*model.User, error)
}

type userRepoImpl struct {
	db *sql.DB
}

func NewRepository() Repository {
	repo := userRepoImpl{}
	db, err := sql.Open("sqlite3", "./aaa.db")
	if err != nil {
		fmt.Printf("connect db fail: %s\n", err)
	}
	repo.db = db
	return &repo
}

var _ Repository = &userRepoImpl{}

func (u *userRepoImpl) GetUserByID(id int64) (*model.User, error) {
	return nil, nil
}
