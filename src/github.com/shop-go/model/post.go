package model

import (
	"time"

	"github.com/hb-go/echo-web/module/log"
)

func (p *Post) GetPostById(id uint64) *Post {
	post := Post{}
	if err := DB().Where("id = ?", id).First(&post).Error; err != nil {
		log.Debugf("Get post error: %v", err)
		return nil
	}

	if err := DB().Model(&post).Related(&post.User).Error; err != nil {
		log.Debugf("Post user related error: %v", err)
		return &post
	}

	return &post
}

func (p *Post) GetUserPostsByUserId(userId uint64, page int, size int) *[]Post {
	posts := []Post{}
	if err := DB().Where("user_id = ?", userId).Offset((page - 1) * size).Limit(size).Find(&posts).Error; err != nil {
		log.Debugf("Get user posts error: %v", err)
		return nil
	}

	for key, post := range posts {
		if err := DB().Model(&post).Related(&post.User).Error; err != nil {
			log.Debugf("Post user related error: %v", err)
		}
		posts[key] = post
	}

	return &posts
}

func (p *Post) PostSave() {
	tx := DB().Begin()

	post1 := Post{Title: "标题3"}
	if err := tx.Create(&post1).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	post2 := Post{Title: "标题4"}
	if err := tx.Create(&post2).Error; err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}

type Post struct {
	Id        uint64    `json:"id,omitempty"`
	UserId    uint64    `form:"user_id" json:"user_id,omitempty"`
	Title     string    `form:"title" json:"title,omitempty"`
	Context   string    `form:"context" json:"context,omitempty"`
	CreatedAt time.Time `gorm:"column:created_time" json:"created_time,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_time" json:"updated_time,omitempty"`

	User User `gorm:"ForeignKey:UserId;AssociationForeignKey:Id" json:"user"`
}

func (p Post) TableName() string {
	return "post"
}
