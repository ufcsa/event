package dao

import (
	"csa-event/models"
	"log"

	"xorm.io/xorm"
)

// PrizeDao maps to prize table
type PrizeDao struct {
	engine *xorm.Engine
}

// NewPrizeDao init a prize dao object
func NewPrizeDao(en *xorm.Engine) *PrizeDao {
	return &PrizeDao{
		engine: en,
	}
}

// Create creates a new prize in the db
func (dao *PrizeDao) Create(data *models.Prize) error {
	_, err := dao.engine.Insert(data)
	return err
}

// Get gets a prize by its id
func (dao *PrizeDao) Get(id int) *models.Prize {
	data := &models.Prize{ID: id}
	ok, err := dao.engine.Get(data)
	if ok && err == nil {
		return data
	}
	if err != nil {
		log.Printf("Prize dao Get error: %v\n", err)
	}
	data.ID = 0
	return data
}

// GetAll gets all the prize in the order of its status
func (dao *PrizeDao) GetAll() []models.Prize {
	dataList := make([]models.Prize, 0)
	err := dao.engine.
		Asc("status").
		Find(&dataList)

	if err != nil {
		log.Printf("Prize dao GetAll error: %v\n", err)
	}
	return dataList
}

// CountAll returns the number of type of prizes
func (dao *PrizeDao) CountAll() int64 {
	num, err := dao.engine.Count(&models.Prize{})
	if err != nil {
		return 0
	}
	return num
}

// Update updates the information of a prize
func (dao *PrizeDao) Update(data *models.Prize, cols []string) error {
	_, err := dao.engine.ID(data.ID).MustCols(cols...).Update(data)
	return err
}

// IncrementLeftNum increments the left_num field by one
func (dao *PrizeDao) IncrementLeftNum(id, num int) (int64, error) {
	ret, err := dao.engine.ID(id).
		Incr("left_num", num).
		Update(&models.Prize{ID: id})
	return ret, err
}

// DecrementLeftNum decrements the left_num field by one
func (dao *PrizeDao) DecrementLeftNum(id, num int) (int64, error) {
	ret, err := dao.engine.ID(id).
		Decr("left_num", num).
		Where("left_num>=", num).
		Update(&models.Prize{ID: id})
	return ret, err
}
