package controllers

import (
	models "forum/model"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func EditComHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		println("cookie inexistant")
		NotFound(w, r, http.StatusNotFound)
		return
	}
	uuid := cookie.Value
	idCurrent := models.GetIDFromUUID(uuid)
	parts := strings.Split(r.URL.Path, "/")
	idcom, _ := strconv.Atoi(parts[len(parts)-1])
	user := models.GetUser(idCurrent)
	com := models.GetComment(idcom)
	if com.User != user {
		http.Error(w, "Permission Denied", http.StatusForbidden)
		return
	}

	var data models.DataEditCom
	data.CurrentUser = user
	data.Comment = com
	data.Id = r.FormValue("idpost")
	tmpl, err := template.ParseFiles("./view/editCom.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
