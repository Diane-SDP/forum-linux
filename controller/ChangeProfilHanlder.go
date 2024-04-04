package controllers

import (
	"crypto/md5"
	"encoding/hex"
	models "forum/model"
	"html/template"
	"log"
	"net/http"
)

func ChangeProfilHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/changeprofil" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              
	}
	cookie, err := r.Cookie("user")
	if err != nil {
		log.Printf("err")
		NotFound(w, r, http.StatusNotFound)
		return
	}
	id := models.GetIDFromUUID(cookie.Value)
	newname := r.FormValue("changename")
	var exists bool
	erreur := models.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", newname).Scan(&exists)
	if erreur != nil {
		panic(erreur)
	}
	if !exists {
		if newname != "" {
			_, err = models.DB.Exec("UPDATE users Set username = ? Where id = ?", newname, id)
			if err != nil {
				panic(err)
			}
		}
	}

	newpwd := r.FormValue("changepsd")
	currentPWD := r.FormValue("psd")

	prbl := models.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE password = ?)", newpwd).Scan(&exists)
	if prbl != nil {
		panic(prbl)
	}

	var info models.User
	info = models.GetUser(id)
	if currentPWD != "" {
		h := md5.New()
		h.Write([]byte(currentPWD))
		currentPWD = hex.EncodeToString(h.Sum(nil))

		if currentPWD == info.Password {
			if !exists {
				if newpwd != "" {
					h := md5.New()
					h.Write([]byte(newpwd))
					newpwd = hex.EncodeToString(h.Sum(nil))
					_, err = models.DB.Exec("UPDATE users set password = ? where id = ?", newpwd, id)
					if err != nil {
						panic(err)
					}
				}
			}
			http.Redirect(w, r, "/changeprofil?check=passwordchange", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/changeprofil?error=badpassword", http.StatusSeeOther)
		}
	}

	bio := r.FormValue("bio")
	if bio != "" {
		_, err = models.DB.Exec("UPDATE users SET bio = ? WHERE id = ?", bio, id)
		if err != nil {
			panic(err)
		}
	}

	file, fileheader, _ := r.FormFile("img")
	var path string
	if file != nil {
		path = models.Upload(file, fileheader)
		_, err = models.DB.Exec("UPDATE users SET image = ? WHERE id = ?", path, id)
		if err != nil {
			print(err.Error())
		}
	}

	store := r.FormValue("photo")
	if store != "" {
		_, err = models.DB.Exec("UPDATE users SET image = ? WHERE id = ?", store, id)
		if err != nil {
			print(err.Error())
		}
	}
	var isGoogle int
	rows, err := models.DB.Query("SELECT googleAccount FROM users WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Erreur lors de la requête", http.StatusInternalServerError)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&isGoogle)
		if err != nil {
			http.Error(w, "Erreur lors de la connexion", http.StatusInternalServerError)
			panic(err)
		}
	}

	info = models.GetUser(id)
	if isGoogle == 1 {
		info.IsGoogle = true
	}
	tmpl, err := template.ParseFiles("./view/changeprofil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
