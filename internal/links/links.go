package links

import (
	"log"

	database "github.com/Roshantwanabasu/news-clone/internal/pkg/db/migrations/mysql"
	"github.com/Roshantwanabasu/news-clone/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select id, title, address from  Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var links []Link
	for row.Next() {
		var link Link
		err := row.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}

	if err = row.Err(); err != nil {
		log.Fatal(err)
	}
	return links

}
