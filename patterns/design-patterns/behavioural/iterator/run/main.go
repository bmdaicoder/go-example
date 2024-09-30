package main

import (
	"fmt"
	"go-example/patterns/design-patterns/behavioural/iterator"
)

func main() {
	user1 := iterator.User{
		Name: "Dai",
		Age:  26,
	}
	user2 := iterator.User{
		Name: "Tai",
		Age:  23,
	}
	userCollection := &iterator.UserCollection{
		Users: []*iterator.User{&user1, &user2},
	}

	iterator := userCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
