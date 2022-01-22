package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"student-management/api"
	"student-management/store"
	"student-management/utils"
)

func main() {

	token, err := utils.CreateToken(1)
	if err != nil {
		log.Fatal("create token error ", err)
	}

	fmt.Printf("token is %s\n", token)

	claim, err := utils.ValidateToken(token)
	if err != nil {
		log.Fatal("validate token error ", err)
	}

	studentId := claim[utils.STUDENT_ID]
	fmt.Printf("student id is %v\n", studentId)

	//database, err := sql.Open("postgres", "postgresql://localhost:5432/school?sslmode=disable")
	//if err != nil {
	//	log.Fatal("db open error", err)
	//	return
	//}
	//
	//err = database.Ping()
	//if err != nil {
	//	log.Fatal("db error", err)
	//	return
	//}
	//
	//queries := db.New(database)
	//stuParam := db.CreateStudentParams{Name: "Kyoko Koyama", Nrc: "elskdjfi", Age: 23, Gender: "female", Password: "root2"}
	//err = queries.CreateStudent(context.Background(), stuParam)
	//if err != nil {
	//	log.Fatal("create student error", err)
	//}
	//
	//result, err := queries.ListStudents(context.Background())
	//if err != nil {
	//	log.Fatal("list student error", err)
	//}
	//fmt.Println(result)
	//
	store := store.NewStore()
	server := api.NewServer(store)

	if err := server.Start(); err != nil {
		log.Fatal("cannot start")
	}
}
