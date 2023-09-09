package mflix

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongoPractice/modles"
	"mongoPractice/mongoClient"
)

type DB struct {
	Name string
}

func (db DB) GetCollection(collectionName string) *mongo.Collection {
	return mongoClient.MongoClient.Database(db.Name).Collection(collectionName)
}


// GetMovieWithIMDB to get the movie with this much imdb and rated by certain no of people
func (db DB) GetMovieWithIMDB(imdb float32,votes,page,limit int) (*[]modles.Movie, error) {
	coll := db.GetCollection("movies")
	// applying and condition on rating and movie
	filter := bson.D{
		{"imdb.rating", bson.D{{"$gt",imdb}}},
		{"imdb.votes",bson.D{{"$gt",votes}}},
	}

	option := options.FindOneOptions{}
	option.SetSkip(int64((page - 1) * limit))

	cursor, err := coll.Find(context.TODO(), filter)

	// un-marshaling data into array of restaurant
	var movies []modles.Movie
	err = cursor.All(context.TODO(), &movies)
	if err != nil {
		return nil, err
	}

	return &movies, nil
}

func (db DB)GetMovieWithCast(casts []string)(*[]modles.Movie,error)  {
	coll := db.GetCollection("movies")
	casts = []string{"Pearl White", "Crane Wilbur","Leo Willis","Anne Hathaway"}
    castArray := bson.A{}
	for _,cast := range casts{
		castArray = append(castArray, cast)
	}

	// The movie having at least one of the cast
	filter := bson.D{{"cast",bson.D{
		{"$in",
		castArray,
		},
	}}}

	cursor, err := coll.Find(context.TODO(), filter)

	// un-marshaling data into array of restaurant
	var movies []modles.Movie
	err = cursor.All(context.TODO(), &movies)
	if err != nil {
		return nil, err
	}

	return &movies, nil
}