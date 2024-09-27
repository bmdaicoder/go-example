package main

import "go-example/patterns/structural/bridge"

func main() {
	url := "127.0.0.1"
	conn := bridge.NewMongodbConnection(url)
	db := bridge.NewMongoDatabase(conn)

	repo := bridge.NewUserRepository(db)
	repo.Create(bridge.User{
		FirstName: "Bui",
		LastName:  "Dai",
		Email:     "bmdaicoder@gmail.com",
	})

	repo.Find("1")
}
