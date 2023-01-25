package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zhaoziliang2019/blog-service/pkg/app"
)

type TagSwagger struct {
	List   []*Tag
	Pagern *app.Pager
}

func (t Tag) Get(db *gorm.DB) (*Tag, error) {
	var tag *Tag
	db = db.Where("state=?", t.State)
	db = db.Where("id =?", t.ID)
	if err := db.Model(&t).Where("is_del=?", 0).First(tag).Error; err != nil {
		return tag, err
	}
	return tag, nil
}
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State)
	if err := db.Model(&t).Where("is_del=?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}
	db = db.Where("state=?", t.State)
	if err = db.Where("is_del=?").Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Updates(values).Where("id=? and is_del=?", t.ID).Error
}
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del=?", t.Model.ID, 0).Delete(&t).Error
}
