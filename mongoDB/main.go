package main

import "mongoDB_trening/mongoDB/router"

func main() {
	db := router.Connect()
	db.DataBase("test")
	db.Collection("test")
	db.Put(30)
	db.DeleteAll()
	db.DeleteByName("Adam")
	db.FindByName("Kelvin")
}
