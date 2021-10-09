package dao

import (
	"log"

	. "github.com/mlabouardy/movies-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)
const (
	COLLECTIONusers = "users"
)
const (
	COLLECTIONposts = "posts"
)

// Establish a connection to database
func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *MoviesDAO) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}
func (m *MoviesDAO) FindAllUsers() ([]User, error) {
	var users1 []User
	err := db.C(COLLECTIONusers).Find(bson.M{}).All(&users1)
	return users1, err
}

// Find a movie by its id
func (m *MoviesDAO) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}
func (m *MoviesDAO) FindUserById(id string) (User, error) {
	var user1 User
	err := db.C(COLLECTIONusers).FindId(bson.ObjectIdHex(id)).One(&user1)
	return user1, err
}
func (m *MoviesDAO) FindPostById(id string) (Posts, error) {
	var post1 Posts
	err := db.C(COLLECTIONposts).FindId(bson.ObjectIdHex(id)).One(&post1)
	return post1, err
}
func (m *MoviesDAO) FindPostsbyUser(id string) ([]Posts, error) {
	var post []Posts
	err := db.C(COLLECTIONposts).Find(bson.M{"user": id}).All(&post)
	return post, err
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}
func (m *MoviesDAO) Insertuser(user1 User) error {
	err := db.C(COLLECTIONusers).Insert(&user1)
	return err
}
func (m *MoviesDAO) Insertpost(post1 Posts) error {
	err := db.C(COLLECTIONposts).Insert(&post1)
	return err
}

// Delete an existing movie
func (m *MoviesDAO) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update an existing movie
func (m *MoviesDAO) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
