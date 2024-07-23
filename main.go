package main

import (
	"log"
	"redis/db"
	"redis/usecase"
)

func main() {
	db, err := db.ConnectToRedis()

	if err != nil {
		log.Println("error in connecting to redis ", err)
		return
	}
	redisRepo := usecase.NewRedisUseCase(db)
	redisRepo.Set("name", "Navneet Shukla")
	redisRepo.Get("name")

	redisRepo.ListPush("player", "Suarez", false)
	redisRepo.ListPush("player", "Messi", false)
	redisRepo.ListPush("player", "Rohit Sharma", true)
	redisRepo.ListPush("player", "MS Dhoni", true)
	redisRepo.ListPop("player", false)
	redisRepo.ListPop("player", true)

}
