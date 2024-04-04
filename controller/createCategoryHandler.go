package controllers

import (
	models "forum/model"
	"html/template"
	"net/http"
)

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createcategory" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return
	}
	tmpl, err := template.ParseFiles("./view/createCategory.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data models.DataProfil
	cookie, err := r.Cookie("user")
	if err != nil {
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return
	}

	id := models.GetIDFromUUID(cookie.Value)
	data.User = models.GetUser(id)

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
