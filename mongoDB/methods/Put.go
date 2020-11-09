package methods

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"sync"
	"time"

	"mongoDB_trening/mongoDB/models"
	"mongoDB_trening/mongoDB/random"
)


func One(collection *mongo.Collection)  {
	p := models.People{
		Name: random.Name(rand.Intn(2738)),
		Age: rand.Intn(80),
		Gender: random.Gender(rand.Intn(2)),
		Country: random.Country(rand.Intn(193)),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	insert, err := collection.InsertOne(ctx, p)
	fmt.Println(insert)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Inserted single documents: ", insert.InsertedID)


}

func Multy(count int, collection *mongo.Collection)  {
	peopple := make([]interface{}, 0, count)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			mu.Lock()
			p := models.People{
				Name: random.Name(rand.Intn(2738)),
				Age: rand.Intn(80),
				Gender: random.Gender(rand.Intn(1)),
				Country: random.Country(rand.Intn(193)),
			}
			defer wg.Done()
			peopple = append(peopple, p)
			mu.Unlock()
		}()
	}
	wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	_, err := collection.InsertMany(ctx, peopple)
	if err != nil{
		log.Println(err)
	}
	fmt.Println("Inserted multiple documents")

}
