module b_good

go 1.15

require (
	be_good/src/db v0.0.0
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/lib/pq v1.8.0 // indirect
)

replace be_good/src/db => ./src/db
