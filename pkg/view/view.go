package view

import (
	"html/template"
	"io"
	"jasy/goblog/pkg/logger"
	"jasy/goblog/pkg/route"
	"path/filepath"
	"strings"
)

func Render(w io.Writer, name string, data interface{}) {
	viewDir := "resources/views/"

	name = strings.Replace(name, ".", "/", -1)

	// 3 所有布局模板文件 Slice
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	// 4 在 Slice 里新增我们的目标文件
	newFiles := append(files, viewDir+name+".gohtml")

	// 5 解析所有模板文件
	tmpl, err := template.New(name + ".gohtml").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(newFiles...)
	logger.LogError(err)

	// 6 渲染模板
	tmpl.ExecuteTemplate(w, "app", data)
}