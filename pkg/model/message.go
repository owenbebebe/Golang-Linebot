package model

import (
	"context"
	"fmt"
	"time"

	"github.com/owenbebebe/Golang-Linebot/pkg/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	MessageId string             `bson:"messageId" json:"messageId"`
	Message   string             `bson:"message" json:"message"`
	UserId    string             `bson:"userId" json:"userId"`
	UserName  string             `bson:"userName" json:"userName"`
}

// creates a message struct in the database
func CreateMessage(m *Message) error {
	m.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Collection.InsertOne(ctx, m)
	//print a string
	fmt.Println("Successfully created message in database")
	return err
}
