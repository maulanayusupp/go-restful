package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rendyfebry/go-restful/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetPersons get all persons
func GetPersons(w http.ResponseWriter, r *http.Request) {
	session := utils.GetMongoSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var errDB error
	c := session.DB("test_db").C("persons")

	results := make([]interface{}, 0)
	errDB = c.Find(bson.M{}).Select(bson.M{"_id": 0}).Sort("-name").All(&results)

	if errDB != nil {
		panic(errDB)
	}

	utils.SendJSONResponse(w, 0, "Success", results)
}

// GetPersonByEmail search person by email
func GetPersonByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	session := utils.GetMongoSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var errDB error
	c := session.DB("test_db").C("persons")

	var result interface{}
	errDB = c.Find(bson.M{"email": email}).Select(bson.M{"_id": 0}).Sort("-name").One(&result)

	if errDB != nil {
		if errDB.Error() == "not found" {
			utils.SendJSONResponse(w, 404, "Not Found", nil)
			return
		}

		panic(errDB)
	}

	utils.SendJSONResponse(w, 0, "Success", result)
}