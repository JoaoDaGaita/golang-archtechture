package main

import (
	"fmt"

	architecture "github.com/GoesToEleven/golang-architecture"
	"github.com/GoesToEleven/golang-architecture/cmd/storage/harddrive"
	"github.com/GoesToEleven/golang-architecture/cmd/storage/mongo"
)

func main() {
	storage := harddrive.Db{}

	u1 := architecture.User{
		First: "James",
	}

	u2 := architecture.User{
		First: "Jenny",
	}

	architecture.Put(storage, 1, u1)
	architecture.Put(storage, 2, u2)

	fmt.Println(architecture.Get(storage, 1))
	fmt.Println(architecture.Get(storage, 2))

	storage2 := mongo.Db{}

	u3 := architecture.User{
		First: "James",
	}

	u4 := architecture.User{
		First: "Jenny",
	}

	architecture.Put(storage2, 1, u3)
	architecture.Put(storage2, 2, u4)

	fmt.Println(architecture.Get(storage2, 1))
	fmt.Println(architecture.Get(storage2, 2))
}
