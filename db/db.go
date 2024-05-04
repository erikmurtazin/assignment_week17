package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type DbRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type DbResponse struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

func NewStorage() (*Mongodb, error) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()
	options := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	collection := client.Database("getircase-study").Collection("records")
	return &Mongodb{
		Client:     client,
		Collection: collection,
	}, nil

}

func (db *Mongodb) FetchDataFromMongo(r DbRequest) (*[]DbResponse, error) {
	startDate, err := parseTime(r.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := parseTime(r.EndDate)
	if err != nil {
		return nil, err
	}
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*10)
	defer cancle()
	pipeLine := bson.A{
		bson.M{
			"$project": bson.M{
				"key":       1,
				"createdAt": 1,
				"totalCount": bson.M{
					"$sum": "$counts",
				},
				"_id": 0,
			},
		},
		bson.M{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": r.MinCount,
					"$lt": r.MaxCount,
				},
				"createdAt": bson.M{
					"$gte": primitive.NewDateTimeFromTime(startDate),
					"$lte": primitive.NewDateTimeFromTime(endDate),
				},
			},
		},
	}
	cursor, err := db.Collection.Aggregate(ctx, pipeLine)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var results []DbResponse
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}
	return &results, nil
}

func parseTime(s string) (time.Time, error) {
	d, err := time.Parse("2006-01-02", s)
	return d, err
}
