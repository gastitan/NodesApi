package service

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"gopkg.in/mgo.v2/bson"
	"gotest.com/nodes-api/database"
	"gotest.com/nodes-api/model"
	"gotest.com/nodes-api/utils"
)

var ctx = context.TODO()

type NodeService interface {
	Save(model.Node) (model.Node, error)
	FindAll() ([]*model.Node, error)
	FindById(id int) (model.Node, error)
	Update(nodeToUpdate model.Node) (model.Node, error)
	Nearest(lat float64, lng float64) (model.Nearest, error)
}

func Save(store model.Node) (model.Node, error) {
	collection := database.GetCollection()

	existentNode, _ := FindById(store.ID)
	if existentNode.ID != 0 {
		return store, errors.New("id already in use")
	}

	_, err := collection.InsertOne(ctx, store)
	if err != nil {
		log.Printf("Could not create Node: %v", err)
		return store, err
	}

	return store, nil
}

func FindAll() ([]model.Node, error) {

	collection := database.GetCollection()
	nodes := []model.Node{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Get Cursor
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Error en db.Find  %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Get all values
	err = cursor.All(ctx, &nodes)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}

	return nodes, nil
}

func FindById(id int) (model.Node, error) {
	collection := database.GetCollection()
	node := model.Node{}

	collection.FindOne(ctx, bson.M{"id": id}).Decode(&node)

	return node, nil
}

func Update(nodeToUpdate model.Node) (model.Node, error) {
	collection := database.GetCollection()
	updatedNode := model.Node{}

	filter := bson.D{primitive.E{Key: "id", Value: nodeToUpdate.ID}}

	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "location", Value: nodeToUpdate.Location},
			primitive.E{Key: "address", Value: nodeToUpdate.Address},
			primitive.E{Key: "nodeType", Value: nodeToUpdate.NodeType},
			primitive.E{Key: "businessHour", Value: nodeToUpdate.BusinessHour},
			primitive.E{Key: "capacity", Value: nodeToUpdate.Capacity},
		}},
	}

	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedNode)
	if err != nil {
		log.Fatal(err)
	}

	return nodeToUpdate, nil
}

func Nearest(lat float64, lng float64) (model.Nearest, error) {
	nodes, err := FindAll()
	if err != nil {
		return model.Nearest{}, err
	}

	locationRequest := model.Location{Lat: lat, Lng: lng}

	minDistance := utils.Distance(locationRequest, nodes[0].Location)
	node := nodes[0]

	for _, s := range nodes {
		distance := utils.Distance(locationRequest, s.Location)
		if distance < minDistance {
			minDistance = distance
			node = s
		}
	}

	log.Printf("The Node %d is the closest for requested distance %v", node.ID, locationRequest)
	return model.Nearest{Distance: minDistance, Node: node}, nil
}
