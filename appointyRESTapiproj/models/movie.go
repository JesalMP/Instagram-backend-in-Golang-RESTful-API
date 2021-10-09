package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document

type Posts struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Captions  string        `json:"captions,omitempty" bson:"captions,omitempty"`
	Imageurl  string        `json:"imageurl" bson:"imageurl,omitempty"`
	Timestamp string        `json:"timestamp" bson:"timestamp,omitempty"`
	UserDat   string        `json:"user" bson:"user"`
}
type User struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name,omitempty" bson:"name,omitempty"`
	Email string        `json:"email" bson:"email,omitempty"`
	Pass  string        `json:"pass" bson:"pass,omitempty"`
}
