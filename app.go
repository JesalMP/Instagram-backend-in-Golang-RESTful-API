package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"gopkg.in/mgo.v2/bson"

	. "github.com/JesalMP/appointyRESTproj/config"
	. "github.com/JesalMP/appointyRESTproj/dao"
	. "github.com/JesalMP/appointyRESTproj/models"
)

var config = Config{}
var dao = MoviesDAO{}
var (
	mutex sync.Mutex
)
var (
	getUserRe   = regexp.MustCompile(`^\/users\/([a-zA-Z0-9]+)$`)
	getPostRe   = regexp.MustCompile(`^\/posts\/([a-zA-Z0-9]+)$`)
	getUserPost = regexp.MustCompile(`^\/posts\/users\/([a-zA-Z0-9]+)$`)
)

func AllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	users1, err := dao.FindAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users1)
}

//
//
func FinduserEndpoint(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	user, err := dao.FindUserById(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	mutex.Unlock()
	respondWithJson(w, http.StatusOK, user)
}
func FindPostEndpoint(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	post, err := dao.FindPostById(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Post ID")
		return
	}
	mutex.Unlock()
	respondWithJson(w, http.StatusOK, post)
}
func FindPostsbyUserEndpoint(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	id := strings.TrimPrefix(r.URL.Path, "/posts/users/")
	post, err := dao.FindPostsbyUser(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	mutex.Unlock()
	respondWithJson(w, http.StatusOK, post)
}

//
//

func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer r.Body.Close()
	var curuser User
	if err := json.NewDecoder(r.Body).Decode(&curuser); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	curuser.ID = bson.NewObjectId()
	if err := dao.Insertuser(curuser); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	mutex.Unlock()
	respondWithJson(w, http.StatusCreated, curuser)
}
func CreatePostEndPoint(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer r.Body.Close()
	var curpost Posts
	if err := json.NewDecoder(r.Body).Decode(&curpost); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	curpost.ID = bson.NewObjectId()
	if err := dao.Insertpost(curpost); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	mutex.Unlock()
	respondWithJson(w, http.StatusCreated, curpost)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	mux := http.NewServeMux()
	//r := mux.NewRouter() //
	mux.HandleFunc("/user/", FinduserEndpoint)
	mux.HandleFunc("/user", CreateUserEndPoint)
	mux.HandleFunc("/posts", CreatePostEndPoint)
	mux.HandleFunc("/posts/", FindPostEndpoint)
	mux.HandleFunc("/posts/users/", FindPostsbyUserEndpoint)
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
