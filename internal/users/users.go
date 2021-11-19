package users

import (
	"database/sql"
	"log"

	database "github.com/Roshantwanabasu/gql-in-go-demo/internal/pkg/db/migrations/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	stmt, err := database.Db.Prepare("Insert into Users(username,password) values (?,?)")
	print(stmt)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

}

func GetUserIdByUsername(username string) (int, error) {
	stmt, err := database.Db.Prepare("Select ID from Users where  Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	return Id, err
}

func (user *User) Authenticate() bool {
	stmt, err := database.Db.Prepare("Select Password from Users where username = ?")
	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRow(user.Username)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)

		}

	}

	return CheckHashedPassword(user.Password, hashedPassword)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
