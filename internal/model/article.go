package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zhaoziliang2019/blog-service/pkg/app"
)

type ArticleSwagger struct {
	List  []*Article
	Paper *app.Pager
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Updates(values).Where("id=? and is_del=?", a.ID).Error; err != nil {
		return err
	}
	return nil
}
func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id=? and state=? and is_del=?", a.ID, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id =? and is_del=?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

type ArticleRow struct {
	ArticleID      uint32
	TagID          uint32
	TagName        string
	ArticleTitle   string
	ArticleDesc    string
	ConverImageUrl string
	Content        string
}

func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id as article_id", "ar.title as article_title", "ar.desc as article_desc", "ar.cover_image_url",
		"ar.content"}
	fields = append(fields, []string{"t.id as tag_id", "t.name as tag_name"}...)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+"as at").
		Joins("Left JOIN `"+Tag{}.TableName()+"` as t on at.tag_id=t.id").
		Joins("left join `"+Article{}.TableName()+"` as ar on at.article_id=ar.id").
		Where("at.`tag_id`=? and ar.state=? and ar.is_del=?", tagID, a.State, 0).Rows()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.ConverImageUrl, &r.Content,
			&r.TagID, &r.TagName); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}
func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" as at").
		Joins("left join `"+Tag{}.TableName()+"` as t on at.tag_id=t.id").
		Joins("left join `"+Article{}.TableName()+"` as ar on at.article_id=ar.id").
		Where("at.`tag_id`=? and ar.state=? and ar.is_del=?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (a ArticleTag) GetByID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id=? and is_del=?", a.ArticleID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return ArticleTag{}, err
	}
	return articleTag, nil
}
func (a ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id=? and is_del=?", a.TagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}
	return articleTags, nil
}
func (a ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?) and is_del=?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}
func (a ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("article_id=? and is_del=?", a.ArticleID, 0).Limit(1).Updates(values).Error; err != nil {
		return err
	}
	return nil
}
func (a ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Where("id=? and is_del=?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}
func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	if err := db.Where("article_id=? and is_del=?", a.ArticleID, 0).Delete(&a).Limit(1).Error; err != nil {
		return err
	}
	return nil
}
