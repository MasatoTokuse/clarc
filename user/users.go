package users

import (
	"context"
	"database/sql"

	"github.com/mtoku/di/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewUserRepository(context context.Context, db *sql.DB) UserRepository {
	return UserRepository{
		DB:      db,
		Context: context,
	}
}

type UserRepository struct {
	DB      *sql.DB
	Context context.Context
}

func (repo UserRepository) Insert() (models.User, error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		UserID:   "mst11",
		Password: "pass",
		Nickname: "tks",
	}

	err = user.Insert(repo.Context, tx, boil.Infer())
	if err != nil {
		return models.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return models.User{}, err
	}

	return user, nil
}
