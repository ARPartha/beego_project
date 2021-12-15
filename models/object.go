package models

import (
	"errors"
	"database/sql"

	"fmt"

_ 	"github.com/lib/pq"
)
var (
	Objects map[string]*Object
)


const (

	host = "localhost"
	
	port = 5432
	
	user = "postgres"
	
	password = "newPassword"
	
	dbname = "postgres"
	
	)


type Object struct {
	FirstName string `json:"FirstName"`
    LastName  string `json:"LastName"`
    Phone     string `json:"Phone"`
    Email     string `json:"Email"`
    Password  string `json:"Password"`
    DoB 	string `json:"DoB"`
}
func init() {
	Objects = make(map[string]*Object)
	//Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	//Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}
}

func AddOne(object Object) (ObjectId string) {

	//asasjkas
	// connection string

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)



	// open database

	db, err := sql.Open("postgres", psqlconn)

	CheckError(err)



	// sqlStatement := `

	// INSERT INTO Person_Information (FirstName, LastName, Email, Phone, Password, DateOfBirth)

	// VALUES ('Shadman', 'Sakib', 'jon@calhoun.io', '01767079174', '1234', '30-08-1997')`

	// _, err = db.Exec(sqlStatement)

	if err != nil {

	panic(err)

	}

	// close database

	defer db.Close()



	// check db

	err = db.Ping()

	CheckError(err)



	fmt.Println("Connected!")
	//ashkas

	Objects[object.Email] = &object
	return object.Email
}

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}

// func Update(ObjectId string, Score int64) (err error) {
// 	if v, ok := Objects[ObjectId]; ok {
// 		v.Score = Score
// 		return nil
// 	}
// 	return errors.New("ObjectId Not Exist")
// }

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}

func CheckError(err error) {

	if err != nil {
	
	panic(err)
	
	}
	
}