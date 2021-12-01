package article

import (
	"jasy/goblog/app/models"
	"jasy/goblog/app/models/user"
	"jasy/goblog/pkg/route"
)

type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   user.User

}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.BaseModel.GetStringID())
}

// CreatedAtDate 创建日期
func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}