package main

import (
	db "Tugas7_ViraMistika/echo-rest/db"
	routes "Tugas7_ViraMistika/echo-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))

}
