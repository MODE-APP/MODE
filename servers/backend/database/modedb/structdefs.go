package modedb

import "go.mongodb.org/mongo-driver/bson/primitive"

//Represents a User in the database
type User struct {
	id,
	profilePicture,
	backgroundPicture primitive.ObjectID

	email,
	displayname,
	username,
	biography string

	followers []primitive.ObjectID
	following []primitive.ObjectID
	posts     []primitive.ObjectID
}

type Post struct {
	id,
	media primitive.ObjectID
	mediaType,
	captionText,
	captionType string
	timeStamp int
	location  interface{} //implement later
	comments  []Comment
}

type Comment struct {
	id,
	author primitive.ObjectID
	text      string
	timeStamp int
	likes     []primitive.ObjectID
	replies   []Comment
}

type Topic struct {
	id,
	featuredImage primitive.ObjectID
	featured  bool
	name      string
	timeStamp int
	followers []primitive.ObjectID
}
