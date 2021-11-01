package article

import (
	"jasy/goblog/pkg/route"
	"strconv"
)

type Article struct {
	ID    int64
	Title string
	Body  string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(a.ID, 10))
}