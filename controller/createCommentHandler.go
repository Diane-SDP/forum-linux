package controllers

import (
	models "forum/model"
	"log"
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
	commentaire := r.FormValue("commentInput")

	_, err = models.DB.Exec("INSERT INTO comment (idUser, idPost, commentaire) VALUES (?, ?, ?)", idUser, idstr, commentaire)
	if err != nil {
		http.Error(w, "Erreur lors de la publication du commentaire", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/commentaire/"+idstr, http.StatusSeeOther)
}
