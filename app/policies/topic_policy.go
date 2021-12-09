package policies

import (
	"jasy/goblog/app/models/article"
	"jasy/goblog/pkg/auth"
)

func CanModifyArticle(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}