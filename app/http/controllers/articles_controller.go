package controllers

import (
	"fmt"
	"gorm.io/gorm"
	"html/template"
	"jasy/goblog/app/models/article"
	"jasy/goblog/pkg/logger"
	"jasy/goblog/pkg/route"
	"jasy/goblog/pkg/types"
	"net/http"
)

type ArticlesController struct {

}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request)  {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)

		tmpl.Execute(w, article)
	}
}
