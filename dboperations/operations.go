package operations

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

type Data struct {
	Id    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Age   int                `bson:"age" json:"age"`
	Login bool               `bson:"login" json:"login"`
}

func (d *Data) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(d)
}

func (d *Data) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(d)
}

func (d *Data) Create() string {

	var Id string
	col := GetDBCollection()
	if col != nil {

		data := &Data{
			Name:  d.Name,
			Age:   d.Age,
			Login: d.Login,
		}
		InsertResult, InserErr := col.InsertOne(context.Background(), data)
		if InserErr != nil {
			log.Printf("%v\n", InserErr)
		}
		Pid, ok := InsertResult.InsertedID.(primitive.ObjectID)
		Id = Pid.Hex()
		if !ok {
			log.Println("Error Obtaining Object ID")
		}
		return Id
	} else {
		return ""
	}

}
func (d *Data) Get() *Data {
	var response = &Data{}
	col := GetDBCollection()
	if col != nil {

		oid := primitive.ObjectID(d.Id)

		filter := bson.M{
			"_id": oid,
		}
		result := col.FindOne(context.Background(), filter)
		if result == nil {
			log.Println("No data is pressent for the given object ID")
			return nil
		}
		decerr := result.Decode(response)

		if decerr != nil {
			log.Printf("Decode Error :%v\n", decerr)
			return nil
		}

	} else {
		return nil
	}

	return response
}

func (d *Data) Update() (old, new *Data) {
	OldDocument := d.Get()
	NewData := &Data{
		Id:    d.Id,
		Name:  d.Name,
		Age:   d.Age,
		Login: d.Login,
	}
	filter := bson.M{
		"_id": d.Id,
	}

	_, err := Collection.ReplaceOne(context.Background(), filter, NewData)
	if err != nil {
		log.Printf("Error while Updating %v\n", err)
	}

	NewDocument := NewData.Get()

	return OldDocument, NewDocument

}

func (d *Data) Delete() (int, bool) {
	response := d.Get()

	if response == nil {
		return 0, false
	} else {
		deleteResult, err := Collection.DeleteOne(context.Background(), bson.M{"_id": response.Id})
		if err != nil {
			log.Printf("Failed Deleting Document :%v\n", err)
		}
		Count := deleteResult.DeletedCount
		return int(Count), true
	}
}

func GetDBCollection() *mongo.Collection {

	if Collection != nil {
		return Collection
	} else {
		log.Println("Creating Mongo DB connection")

		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Println(err)
			return nil
		}
		err = client.Connect(context.TODO())
		if err != nil {
			log.Println(err)
			return nil
		}
		Collection = client.Database("UserDB").Collection("Users")
		return Collection
	}
}
