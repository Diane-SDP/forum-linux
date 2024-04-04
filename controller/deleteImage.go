package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
)

func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	idint, _ := strconv.Atoi(parts[len(parts)-1])
	println(r.URL.Path)
	_, err := models.DB.Exec("UPDATE posts SET image = ? WHERE id = ?", "", idint)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image supprimée avec succès"))
}
