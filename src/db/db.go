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
	Id              string  `json:"Id"`
	Companyname     string  `json:"Cmp_name"`
	Companylocation string  `json:"Cmp_location"`
	Jobtitle        string  `json:"Job_title"`
	Jobdescription  string  `json:"Job_description"`
	Applylink       string  `json:"Apply_link"`
	Timestamps      float32 `json: "Timestamps"`
}

//GetData returns array of Job structs
func GetData(limit int8) []Job {
	//returns data from postgres db
	returnArray := make([]Job, limit)
	count := 0
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	password := os.Getenv("USER_PASS")
	user := os.Getenv("USER_NAME")
	dbname := "career_bot"
	psqlInfo := fmt.Sprintf("host=52.188.71.209 port=5432 user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
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
	userSql := fmt.Sprintf("SELECT * FROM cs_bot.Jobs limit %d", limit)
	rows, err := db.Query(userSql)
	defer rows.Close()
	for rows.Next() {
		var myJob Job
		err := rows.Scan( &myJob.Id,&myJob.Companyname, &myJob.Companylocation, &myJob.Jobtitle, &myJob.Jobdescription, &myJob.Applylink, &myJob.Timestamps)
		if err != nil {
			log.Fatal("Couldn't execute for some odd reason", err)
		}
		returnArray[count] = myJob
		count++

	}
	if err != nil {
		log.Fatal("Couldn't execute for some odd reason", err)
	}
	return returnArray
}
