package models

type Comment struct {
	Content string
	User    User
	Id      int
	IsMine  bool
}

func GetComments(idpost int) []Comment {
	rows, err := DB.Query("SELECT id, idUser, content FROM comment WHERE idPost = ?", idpost)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var comments []Comment
	var iduser int
	for rows.Next() {
		var com Comment
		err := rows.Scan(&com.Id, &iduser, &com.Content)
		if err != nil {
			panic(err)
		}
		user := GetUser(iduser)
		com.User = user
		comments = append(comments, com)
	}
	return comments
}

func GetComment(id int) Comment {
	rows, err := DB.Query("SELECT id, idUser, content FROM comment WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var comment Comment
	var iduser int
	for rows.Next() {
		err := rows.Scan(&comment.Id, &iduser, &comment.Content)
		if err != nil {
			panic(err)
		}
		user := GetUser(iduser)
		comment.User = user
	}
	return comment
}

func DeleteCommentsOf(idpost int) {
	_, err := DB.Exec("DELETE FROM comment WHERE idPost = ?", idpost)
	if err != nil {
		panic(err)
	}
}
