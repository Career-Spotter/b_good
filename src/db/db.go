package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Job struct {
	Id              string `json:"id"`
	JobTitle        string `json:"job_title"`
	ApplyLink       string `json:"apply_link"`
	CompanyName     string `json:"cmp_name"`
	CompanyLocation string `json:"cmp_location"`
	JobDescription  string `json:"job_description"`
}

func GetData() Job {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	password := os.Getenv("USER_PASS")
	user := os.Getenv("USER_NAME")
	dbname := "career_bot"
	psqlInfo := fmt.Sprintf("host=13.92.43.146 port=5432 user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sucessfully connected")
	var myJob Job
	userSql := "SELECT * FROM cs_bot.Jobs"
	err = db.QueryRow(userSql).Scan(&myJob.ApplyLink, &myJob.CompanyName, &myJob.CompanyLocation, &myJob.JobDescription, &myJob.JobTitle, &myJob.Id)
	if err != nil {
		log.Fatal("Couldn't execute for some odd reason", err)
	}
	return myJob
}
