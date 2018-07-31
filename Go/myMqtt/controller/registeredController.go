package controller

import (
	"net/http"
	"fmt"
	"html/template"
)

func RegisteredController(w http.ResponseWriter, r *http.Request){

	fmt.Println("-------------这是Value获取的--------------")
	fmt.Println("用户名: ",r.PostFormValue("name"))
	fmt.Println("密码是: ",r.PostFormValue("pass"))

	//页面跳转
	t, _ := template.ParseFiles("resources/templates/create.html")
	t.Execute(w, nil)
}