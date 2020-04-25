package dataaccessobject

import (
	"context"
	"log"
	"retrck/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DAO declaration
type DAO struct {
	Server   string
	Database string
}

var db *mongo.Database

// COLLECTION declaration
const (
	COLLECTION = "properties"
)

// Connection to the MongoDB
func (d *DAO) Connection() {
	clientOpts := options.Client().ApplyURI(d.Server)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(d.Database)
}


// FindAll list of props
func (d *DAO) FindAll() (props []models.Property, err error) {
	cursor, err := db.Collection(COLLECTION).Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &props); err != nil {
		log.Fatal(err)
	}
	return props, err
}