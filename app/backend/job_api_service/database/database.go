package database

import (
	"context"
	"encoding/json"
	"gitlab/jobs/graph/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb+srv://sami:abdul-sami@cluster0.bvheu7o.mongodb.net/game?retryWrites=true&w=majority"

//var connectionString string = "mongodb+srv://sami:abdul-sami@cluster0.yupnqta.mongodb.net/game?retryWrites=true&w=majority"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Println("Database connection error")
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}

}
func (db *DB) jobsList() []*model.Jobs {
	jobCollec := db.client.Database("mernstack").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var joblist []*model.Jobs
	cursor, err := jobCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &joblist); err != nil {
		panic(err)
	}
	return joblist
}

func (db *DB) job(id string) *model.Jobs {
	jobCollec := db.client.Database("mernstack").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var joblist model.Jobs
	err := jobCollec.FindOne(ctx, filter).Decoode(&joblist)
	if err != nil {
		log.Println("Can't find this job")
		log.Fatal(err)
	}
	return joblist
}

func (db *DB) CreateJob(jobInfo model.CreateJobInput) *model.Jobs {
	jobCollec := db.client.Database("mernstack").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	url := "https://gitlab.com/api/v4/projects/38858455/jobs"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + "glpat-xqs1ieNuK2gLtta4yBs8"

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	json.Unmarshal(body, &jobInfo)

	inserted, err := jobCollec.InsertOne(ctx, jobInfo)
	log.Println(inserted)
	return inserted
}
