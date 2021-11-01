package article

import (
	"jasy/goblog/pkg/route"
	"strconv"
)

type Article struct {
	ID    int
	Title string
	Body  string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
}