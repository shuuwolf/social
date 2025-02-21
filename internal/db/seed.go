package db

import (
	"context"
	"log"
	"social/internal/store"
)

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(1)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: "test222123",
			Email:    "test222@test.com",
			Password: "test222123",
		}
	}

	return users
}
