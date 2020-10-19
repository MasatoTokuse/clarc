package user

import (
	"context"
	"database/sql"

	user_domain_model "github.com/mtoku/clarc/app/domain/models/user"
	"github.com/mtoku/clarc/app/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type IUserRepository interface {
	Save(user user_domain_model.User) (*user_domain_model.User, error)
	FindBy(userID, password, name string) (*user_domain_model.User, error)
	Remove(user user_domain_model.User) (*user_domain_model.User, error)
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

func (repo UserRepository) Save(user user_domain_model.User) (*user_domain_model.User, error) {

	if err := user.Valid(); err != nil {
		return &user_domain_model.User{}, err
	}

	dbUser := &models.User{
		UserID:   user.UserID,
		Password: user.Password,
		Name:     user.Name,
	}

	tx, err := repo.DB.Begin()
	if err != nil {
		return &user_domain_model.User{}, err
	}

	if err := dbUser.Insert(repo.Context, tx, boil.Infer()); err != nil {
		return &user_domain_model.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return &user_domain_model.User{}, err
	}

	user.ID = dbUser.ID

	return &user, nil
}

func (repo UserRepository) FindBy(userID, password, name string) (*user_domain_model.User, error) {

	qm := make([]qm.QueryMod, 0)
	if userID != "" {
		qm = append(qm, models.UserWhere.UserID.EQ(userID))
	}
	if password != "" {
		qm = append(qm, models.UserWhere.Password.EQ(password))
	}
	if name != "" {
		qm = append(qm, models.UserWhere.Name.EQ(name))
	}

	dbUser, err := models.Users(qm...).One(repo.Context, repo.DB)
	if err != nil {
		return &user_domain_model.User{}, err
	}

	user := &user_domain_model.User{
		ID:       dbUser.ID,
		UserID:   dbUser.UserID,
		Password: dbUser.Password,
		Name:     dbUser.Name,
	}

	return user, nil
}

func (repo UserRepository) Remove(user user_domain_model.User) (*user_domain_model.User, error) {

	tx, err := repo.DB.Begin()
	if err != nil {
		return &user_domain_model.User{}, err
	}

	dbUser := &models.User{
		ID:       user.ID,
		UserID:   user.UserID,
		Password: user.Password,
		Name:     user.Name,
	}

	if _, err := dbUser.Delete(repo.Context, tx); err != nil {
		return &user_domain_model.User{}, err
	}

	if err := tx.Commit(); err != nil {
		return &user_domain_model.User{}, err
	}

	return &user, nil
}

func (repo UserRepository) CloseDB() error {
	return repo.DB.Close()
}
