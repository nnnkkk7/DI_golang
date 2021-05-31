package main

import (
	"database/sql"
	"fmt"

	"github.com/nnnkkk7/DI_golang/repository"
	"github.com/nnnkkk7/DI_golang/service"
)

func main() {
	db, err := sql.Open("mysql", "sample")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	repo := repository.NewUserRepository(db)
	sv := service.NewUserService(repo)
	fmt.Println(sv)

}
