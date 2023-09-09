package modles

import "go.mongodb.org/mongo-driver/bson/primitive"

type Restaurant struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Cuisine string             `json:"cuisine" bson:"cuisine"`
}

type Movie struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Title    string             `json:"title" bson:"title"`
}
