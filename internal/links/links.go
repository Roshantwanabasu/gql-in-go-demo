package links

import (
	"log"

	database "github.com/Roshantwanabasu/gql-in-go-demo/internal/pkg/db/migrations/mysql"
	"github.com/Roshantwanabasu/gql-in-go-demo/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
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
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from  Links L inner join Users U on L.UserID = U.ID")
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
	var username string
	var id string
	for row.Next() {
		var link Link
		err := row.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		}
		links = append(links, link)
	}

	if err = row.Err(); err != nil {
		log.Fatal(err)
	}
	return links

}
