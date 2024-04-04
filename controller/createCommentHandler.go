package controllers

import (
	models "forum/model"
	"net/http"
	"strings"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return
	}
	idUser := models.GetIDFromUUID(cookie.Value)
	parts := strings.Split(r.URL.Path, "/")
	idstr := parts[len(parts)-1]
	comment := r.FormValue("commentInput")

	_, err = models.DB.Exec("INSERT INTO comment (idUser, idPost, content) VALUES (?, ?, ?)", idUser, idstr, comment)
	if err != nil {
		http.Error(w, "Erreur lors de la publication du commentaire", http.StatusInternalServerError)
		panic(err)
	}

	http.Redirect(w, r, "/comments/"+idstr, http.StatusSeeOther)
}
