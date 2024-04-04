package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
)

func SaveComHandler(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("contentInput")
	id := r.FormValue("id")
	idpost := r.FormValue("idpost")
	idint, _ := strconv.Atoi(id)
	_, err := models.DB.Exec("UPDATE comment SET content = ? WHERE id = ?", content, idint)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/comments/"+idpost, http.StatusSeeOther)
}
