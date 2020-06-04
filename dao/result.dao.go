package dao

import (
	"csa-event/models"
	"log"

	"xorm.io/xorm"
)

// ResultDao is the dao object mapping to result table
type ResultDao struct {
	engine *xorm.Engine
}

// NewResultDao returns a result dao object with the given xorm engine
func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{
		engine: engine,
	}
}

// Get gets the result info by its record id
func (d *ResultDao) Get(id int) *models.Result {
	data := &models.Result{ID: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	if err != nil {
		log.Printf("Result dao Get() error: %v\n", err)
	}
	data.ID = 0
	return data
}

// GetAll returns all the results
func (d *ResultDao) GetAll() []models.Result {
	dataList := make([]models.Result, 0)
	err := d.engine.Desc("id").Find(&dataList)
	if err != nil {
		log.Printf("Result dao GetAll() error : %v\n", err)
	}
	return dataList
}

// Create generates a new result record when a there is a winner
func (d *ResultDao) Create(data *models.Result) error {
	_, err := d.engine.Insert(data)
	return err
}

// Update updates a record of result by its id
func (d *ResultDao) Update(data *models.Result, cols []string) error {
	_, err := d.engine.ID(data.ID).MustCols(cols...).Update(data)
	return err
}

// SearchByUser gets all the prize record by user
func (d *ResultDao) SearchByUser(userID int) []models.Result {
	dataList := make([]models.Result, 0)
	err := d.engine.
		Where("user_id=?", userID).
		Asc("id").
		Find(&dataList)
	if err != nil {
		log.Printf("Result dao SearchByUser() error: %v\n", err)
	}
	return dataList
}

// SearchByGift gets all the prize records by a given gift ID
func (d *ResultDao) SearchByGift(prizeID int) []models.Result {
	dataList := make([]models.Result, 0)
	err := d.engine.
		Where("prize_id=?", prizeID).
		Asc("id").
		Find(&dataList)
	if err != nil {
		log.Printf("Result dao SearchByUser() error: %v\n", err)
	}
	return dataList
}
