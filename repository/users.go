package repository

import "govault/models"

type UsersRepository interface {
	Save(user *models.User)
	Update(user *models.User)
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Delete(id string) error
}
