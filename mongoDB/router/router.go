package router

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mongoDB_trening/mongoDB/methods"
	"time"
)

type DB struct {
	client *mongo.Client
	database *mongo.Database
	collection *mongo.Collection

}

func Connect() *DB{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return &DB{
		client: client,
	}
}

func (db *DB) DataBase(database string) {
	db.database = db.client.Database(database)
}

func (db *DB) Collection(col string) {
	db.collection = db.database.Collection(col)
}

func (db *DB) Put(count int)  {
	if count == 1 {
		methods.One(db.collection)
	} else {
		methods.Multy(count, db.collection)
	}
}

func (db *DB) DeleteAll(){
	methods.DeleteAll(db.collection)
}

func (db *DB) DeleteByName(name string)  {
	methods.DeleteByName(name, db.collection)
}

func (db *DB) FindByName(name string) {
	methods.FindByName(name, db.collection)
}


