package dao

import (
	"log"

	. "github.com/JesalMP/appointyRESTproj/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

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

func (m *MoviesDAO) FindAllUsers() ([]User, error) {
	var users1 []User
	err := db.C(COLLECTIONusers).Find(bson.M{}).All(&users1)
	return users1, err
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

func (m *MoviesDAO) Insertuser(user1 User) error {
	err := db.C(COLLECTIONusers).Insert(&user1)
	return err
}
func (m *MoviesDAO) Insertpost(post1 Posts) error {
	err := db.C(COLLECTIONposts).Insert(&post1)
	return err
}
