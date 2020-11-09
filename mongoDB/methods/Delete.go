package methods

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"strings"
	"time"
)

func DeleteAll(collection *mongo.Collection) {
	LOOP:
		for{
		fmt.Println("Are you sure about this? Write Y/N : ")
		reader := bufio.NewReader(os.Stdin)
		s, _ := reader.ReadString('\n')
		s = strings.ToLower(s)
		if strings.Contains(s,"y") {
			break LOOP
		} else {
			fmt.Println("The removal was a deviation")
			return
		}

	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	delMale, err := collection.DeleteMany(ctx, bson.M{"gender":"Male"})
	if err != nil {
		log.Fatal(err)
	}
	delFemale, err := collection.DeleteMany(ctx, bson.M{"gender":"Female"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data was deleted from DB: ", delMale.DeletedCount + delFemale.DeletedCount)

}

func DeleteByName(Name string, collection *mongo.Collection)  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	del, err := collection.DeleteMany(ctx, bson.M{"name":Name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data was deleted from DB: ", del.DeletedCount)
}
