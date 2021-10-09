package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"gopkg.in/mgo.v2/bson"

	. "github.com/JesalMP/appointyRESTproj/config"
	. "github.com/JesalMP/appointyRESTproj/dao"
	. "github.com/JesalMP/appointyRESTproj/models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = MoviesDAO{}
var (
	mutex sync.Mutex
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

	params := mux.Vars(r)
	user, err := dao.FindUserById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	mutex.Unlock()
	respondWithJson(w, http.StatusOK, user)
}
func FindPostEndpoint(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	params := mux.Vars(r)
	post, err := dao.FindPostById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Post ID")
		return
	}
	mutex.Unlock()
	respondWithJson(w, http.StatusOK, post)
}
func FindPostsbyUserEndpoint(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	params := mux.Vars(r)
	post, err := dao.FindPostsbyUser(params["id"])
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
	//r := http.NewServeMux()
	r := mux.NewRouter() //
	r.HandleFunc("/user/{id}", FinduserEndpoint).Methods("GET")
	r.HandleFunc("/user", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/posts", CreatePostEndPoint).Methods("POST")
	r.HandleFunc("/posts/{id}", FindPostEndpoint).Methods("GET")
	r.HandleFunc("/posts/users/{id}", FindPostsbyUserEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
