package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokemon/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx = context.Background()

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	c := GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon?offset=0&limit=10")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject)
	for i := 0; i < len(responseObject.Results); i++ {
		// log.Println(responseObject.Results[i].Name)
		// log.Println(responseObject.Results[i].Url)

		colection := c.Database("pokemon").Collection("result")
		_, err = colection.InsertOne(ctx,
			model.Result{responseObject.Results[i].Name, responseObject.Results[i].Url})
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("Insert success!")
	}
	poke := GetAllPokemon(c, bson.M{})
	for _, hero := range poke {
		log.Println(hero.Name, hero.Url)
	}
}

func GetAllPokemon(client *mongo.Client, filter bson.M) []*model.Result {
	var pokemons []*model.Result
	collection := client.Database("pokemon").Collection("result")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var pokemon model.Result
		err = cur.Decode(&pokemon)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		pokemons = append(pokemons, &pokemon)
	}
	return pokemons
}

func InsertNewHero(client *mongo.Client, user model.Users) interface{} {
	var users model.Users
	collection := client.Database("pokemon").Collection("users")
	Input, err := collection.InsertOne(ctx, users)
	if err != nil {
		log.Fatalln("Error Input User", err)
	}
	return Input
}
