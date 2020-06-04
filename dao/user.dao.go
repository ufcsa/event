package dao

import (
	"csa-event/models"
	"log"

	"xorm.io/xorm"
)

// UserDao is a dao struct mapping to user table
type UserDao struct {
	engine *xorm.Engine
}

// NewUserDao returns a UserDao object with the given Dao engine
func NewUserDao(en *xorm.Engine) *UserDao {
	return &UserDao{
		engine: en,
	}
}

// Get gets the user's information
func (d *UserDao) Get(id int) *models.User {
	data := &models.User{ID: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	if err != nil {
		log.Printf("User dao Get() error: %v\n", err)
	}
	data.ID = 0
	return data
}

// GetAll gets all the users sorted by id
func (d *UserDao) GetAll() []models.User {
	dataList := make([]models.User, 0)
	err := d.engine.
		Asc("id").
		Find(&dataList)
	if err != nil {
		log.Printf("User dao GetAll() error: %v\n", err)
	}
	return dataList
}

// Create creates a new user
func (d *UserDao) Create(data *models.User) error {
	_, err := d.engine.Insert(data)
	return err
}

// Update updates user's information
func (d *UserDao) Update(data *models.User, columns []string) error {
	_, err := d.engine.ID(data.ID).MustCols(columns...).Update(data)
	return err
}
