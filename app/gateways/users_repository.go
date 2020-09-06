package gateways

import (
	"context"
	"database/sql"

	"github.com/mtoku/di/app/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type IUserRepository interface {
	Save(user models.User) (*models.User, error)
	FindBy(userID, password, nickname string) (*models.User, error)
	Remove(user models.User) (*models.User, error)
	CloseDB() error
}

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

func (repo UserRepository) Save(user models.User) (*models.User, error) {

	if err := user.Valid(); err != nil {
		return &models.User{}, err
	}

	tx, err := repo.DB.Begin()
	if err != nil {
		return &models.User{}, err
	}

	if err := user.Insert(repo.Context, tx, boil.Infer()); err != nil {
		return &models.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (repo UserRepository) FindBy(userID, password, nickname string) (*models.User, error) {

	qm := make([]qm.QueryMod, 0)
	if userID != "" {
		qm = append(qm, models.UserWhere.UserID.EQ(userID))
	}
	if password != "" {
		qm = append(qm, models.UserWhere.Password.EQ(password))
	}
	if nickname != "" {
		qm = append(qm, models.UserWhere.Nickname.EQ(nickname))
	}

	user, err := models.Users(qm...).One(repo.Context, repo.DB)
	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (repo UserRepository) Remove(user models.User) (*models.User, error) {

	tx, err := repo.DB.Begin()
	if err != nil {
		return &models.User{}, err
	}

	if _, err := user.Delete(repo.Context, tx); err != nil {
		return &models.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (repo UserRepository) CloseDB() error {
	return repo.DB.Close()
}
