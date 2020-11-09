package methods

import (
	"context"
	"fmt"
	"log"
	"mongoDB_trening/mongoDB/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindByName(name string, collection *mongo.Collection)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	people := []models.People{}
	poepleByName, err := collection.Find(ctx, bson.M{"name": name})
	for poepleByName.Next(ctx){
		p := models.People{}
		if err = poepleByName.Decode(&p); err != nil {
			log.Fatal(err)
		}
		people = append(people, p)
	}
	if len(people) > 0{
		fmt.Printf("%s was found, count: %d \n ", name, len(people))
		fmt.Println(people)
	} else {
		fmt.Println("No found")
	}

}
