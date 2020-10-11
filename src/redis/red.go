package p_redis

import (
	"be_good/src/db"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//EstablishConn returns a redis client
func EstablishConn() redis.Client {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't get the env file for some odd reason")
	}
	//host  must be ip:port_number
	hostIp := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_PASS")
	rDb := redis.NewClient(&redis.Options{
		Addr:     hostIp,
		Password: redisPass,
		DB:       0, //default database available
	})
	err := rdb.Set(ctx, "test", "1", 0).Err()
	if err != nil {
		panic("couldn't establish connection with redis server ::")
	}
	return rDb
}
//GetData returns an array of Job Struct from redis server
func GetDataFromRedis(limit int) db.Job[]{
	var payLoad[limit] db.Job
	rDb := EstablishConn()
	
}
