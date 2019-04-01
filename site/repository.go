package site

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Repository ...
type Repository struct {
	s *mgo.Session
}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "health-check"

// DOCNAME the name of the document
const DOCNAME = "sites"

// GetSites returns the list of Sites
func (r Repository) GetSites() Sites {
	c := r.s.DB(DBNAME).C(DOCNAME)
	results := Sites{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddSite inserts an Site in the DB
func (r Repository) AddSite(site Site) bool {

	site.ID = bson.NewObjectId()
	site.CreatedAt = time.Now()
	site.UpdatedAt = time.Now()
	err := r.s.DB(DBNAME).C(DOCNAME).Insert(site)

	fmt.Printf("%v", site)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateSite updates an Site in the DB
func (r Repository) PatchSite(match map[string]interface{}, update map[string]interface{}) bool {

	err := r.s.DB(DBNAME).C(DOCNAME).Update(match, update)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteSite deletes an Site
func (r Repository) DeleteSite(id string) string {

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "404"
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := r.s.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "500"
	}

	// Write status
	return "200"
}
