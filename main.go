package main

import (
	"database/sql"
	controllers "forum/controller"
	models "forum/model"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.DB, _ = sql.Open("sqlite3", "database.db")
	defer models.DB.Close()
	_, err := models.DB.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT,
        password TEXT DEFAULT '',
		bio TEXT DEFAULT '',
		image TEXT DEFAULT './Assets/img/user.png',
        email TEXT,
		uuid TEXT,
		googleAccount INTEGER DEFAULT 0
    )`)
	if err != nil {
		panic(err)
	}
	_, err = models.DB.Exec(`CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
		image TEXT,
		likes INTEGER NOT NULL,
		idUser INTEGER,
		idCategory INTEGER,
        FOREIGN KEY (idUser) REFERENCES users(id),
		FOREIGN KEY (idCategory) REFERENCES categories(id)
    )`)
	if err != nil {
		panic(err)
	}
	_, err = models.DB.Exec(`CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
		image TEXT DEFAULT './Assets/img/categories.png',
        description TEXT
    )`)
	if err != nil {
		panic(err)
	}
	_, err = models.DB.Exec(`CREATE TABLE IF NOT EXISTS likes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        idUser INTEGER,
        idPost INTEGER,
		uuidUser TEXT,
		FOREIGN KEY (idUser) REFERENCES user(id),
		FOREIGN KEY (idPost) REFERENCES post(id)
    )`)

	if err != nil {
		panic(err)
	}

	_, err = models.DB.Exec(`CREATE TABLE IF NOT EXISTS comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        idUser INTEGER,
        idPost INTEGER,
		content TEXT
    )`)

	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("Assets"))
	http.Handle("/Assets/", http.StripPrefix("/Assets", fs))

	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/postcategory", controllers.CategoryHandler)
	http.HandleFunc("/create", controllers.CreatePostHandler)
	http.HandleFunc("/like/", controllers.LikedPostHandler)
	http.HandleFunc("/loginsuccess", controllers.LoginSuccessHandler)
	http.HandleFunc("/deco", controllers.DisconnectHandler)
	http.HandleFunc("/topic/", controllers.TopicHandler)
	http.HandleFunc("/edit/", controllers.EditHandler)
	http.HandleFunc("/editCom/", controllers.EditComHandler)
	http.HandleFunc("/createcategory", controllers.CreateCategoryHandler)
	http.HandleFunc("/post", controllers.PostHandler)
	http.HandleFunc("/connection", controllers.ConnectionHandler)
	http.HandleFunc("/signup", controllers.SignupHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/profil/", controllers.ProfilHandler)
	http.HandleFunc("/changeprofil", controllers.ChangeProfilHandler)
	http.HandleFunc("/comments/", controllers.CommentHandler)
	http.HandleFunc("/createcomment/", controllers.CreateCommentHandler)
	http.HandleFunc("/save", controllers.SaveHandler)
	http.HandleFunc("/saveCom", controllers.SaveComHandler)
	http.HandleFunc("/logingoogle", controllers.GoogleLoginHandler)
	http.HandleFunc("/callback", controllers.HandleGoogleCallback)
	http.HandleFunc("/deleteImage/", controllers.DeleteImageHandler)
	http.HandleFunc("/delete", controllers.DeletePostHandler)
	http.HandleFunc("/deleteCom", controllers.DeleteComHandler)

	panic(http.ListenAndServe("localhost:8080", nil))

}
