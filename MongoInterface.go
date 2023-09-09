package main

import "go.mongodb.org/mongo-driver/mongo"

type MongoDB interface {
	GetCollection(collectionName string) *mongo.Collection
}

