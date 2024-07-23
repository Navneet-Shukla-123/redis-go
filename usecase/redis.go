package usecase

import (
	"context"
	"log"
	"redis/db"
)

type RedisUseCase struct {
	DB db.RedisRepo
}

func NewRedisUseCase(redis db.RedisRepo) *RedisUseCase {
	return &RedisUseCase{
		DB: redis,
	}

}

func (r *RedisUseCase) Set(key, value string) {
	ctx := context.Background()
	err := r.DB.SetKey(ctx, key, value)
	if err != nil {
		log.Println("error in setting up the key ", err)
		return
	}

	log.Println("Key is successfully setup")
}

func (r *RedisUseCase) Get(key string) {
	ctx := context.Background()

	value, err := r.DB.GetKey(ctx, key)
	if err != nil {
		log.Println("error in getting up the key ", err)
		return
	}
	log.Printf("Key is %s and value is %s \n", key, value)
}

func (r *RedisUseCase) ListPush(key, value string, side bool) {
	// side is false for left push and side is true for right push in the list

	err := r.DB.ListPush(context.Background(), side, key, value)
	if err != nil {
		if !side {
			log.Println("error in left push of the list ", err)
			return
		} else {
			log.Println("error in right push of the list ", err)
			return
		}
	}
	log.Println("Insertion in list is successfull")

}

func (r *RedisUseCase) ListPop(key string, side bool) {
	// side is false for left push and side is true for right push in the list

	value, err := r.DB.ListPop(context.Background(), side, key)
	if err != nil {
		if !side {
			log.Println("error in popping froom the  left of list ", err)
			return
		} else {
			log.Println("error in popping froom the  right of list ", err)
			return
		}
	}
	if !side {
		log.Println("Value from the left of the list is  ", value)
	} else {
		log.Println("Value from the right of the list is ", value)
	}

}
