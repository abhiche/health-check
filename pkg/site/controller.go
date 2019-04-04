package site

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	Repository Repository
}

type ErrorMessage struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	sites := c.Repository.GetSites() // list of all sites
	log.Println(sites)
	data, _ := json.Marshal(sites)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddSite POST /
func (c *Controller) AddSite(w http.ResponseWriter, r *http.Request) {

	var site Site
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error AddSite", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddSite", err)
	}
	if err := json.Unmarshal(body, &site); err != nil { // unmarshall body contents as a type Site
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddSite unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	print("IsValidRequestURL(site.URL)", IsValidRequestURL(site.URL))
	if IsValidRequestURL(site.URL) == false {
		w.WriteHeader(http.StatusBadRequest)
		data := &ErrorMessage{
			Code: "F1",
			Msg:  "Invalid url received",
		}
		w.Header().Set("Content-Type", "application/json")
		resJSON, err := json.Marshal(data)
		if err != nil {
			log.Fatalln("Error AddSite marshalling response", err)
		}
		w.Write(resJSON)
		return
	}

	success := c.Repository.AddSite(site) // adds the site to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// PatchSite PATCH /
func (c *Controller) PatchSite(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"] // param id

	var sp SitePatch
	sp.UpdatedAt = time.Now()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error AddSite", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddSite", err)
	}
	if err := json.Unmarshal(body, &sp); err != nil { // unmarshall body contents as a type Site
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddSite unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success := c.Repository.PatchSite(bson.M{"uuid": id}, bson.M{"$set": &sp}) // adds the site to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteSite DELETE /
func (c *Controller) DeleteSite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] // param id

	if err := c.Repository.DeleteSite(id); err != "" { // delete a site by id
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	return
}
