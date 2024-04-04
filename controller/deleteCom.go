package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
)

func DeleteComHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idpost := r.FormValue("idpost")
	idint, _ := strconv.Atoi(id)
	_, err := models.DB.Exec("DELETE FROM comment WHERE id = ?", idint)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/comments/"+idpost, http.StatusSeeOther)
}
