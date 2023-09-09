package restaurant

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mongoPractice/modles"
	"mongoPractice/mongoClient"
)

type DB struct {
	Name string
}

func (db DB) GetCollection(collectionName string) *mongo.Collection {
	return mongoClient.MongoClient.Database(db.Name).Collection(collectionName)
}

// GetRestaurantDetail here you can write the quires related to restaurant
func (db DB) GetRestaurantDetail(restName string) (*modles.Restaurant, error) {
	coll := db.GetCollection("restaurants")
	filter := bson.D{{"name", restName}}
	var restInfo modles.Restaurant
	err := coll.FindOne(context.TODO(), filter).Decode(&restInfo)
	if err != nil {
		return nil, err
	}

	return &restInfo, nil
}

// GetAllRestaurantDetail get detail of all restaurant
func (db DB) GetAllRestaurantDetail(cuisine string) (*[]modles.Restaurant, error) {
	coll := db.GetCollection("restaurants")
	// query to get all restaurant
	filter := bson.D{}
	if cuisine != "" {
		filter = bson.D{{"cuisine", cuisine}}
	}

	var restaurants []modles.Restaurant
	cursor, err := coll.Find(context.TODO(), filter)

	// un-marshaling data into array of restaurant
	err = cursor.All(context.TODO(), &restaurants)
	if err != nil {
		return nil, err
	}

	return &restaurants, nil
}
