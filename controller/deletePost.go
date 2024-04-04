package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
)

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	idint, _ := strconv.Atoi(id)
	_, err := models.DB.Exec("DELETE FROM posts WHERE id = ?", idint)
	if err != nil {
		panic(err)
	}

	models.DeleteCommentsOf(idint)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
