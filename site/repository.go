package site

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

//Repository ...
type Repository struct {
	s *mgo.Session
}

// DBNAME the name of the DB instance
const DBNAME = "health-check"

// CNAME the name of the collection
const CNAME = "sites"

// GetSites returns the list of Sites
func (r Repository) GetSites() Sites {
	c := r.s.DB(DBNAME).C(CNAME)
	results := Sites{}
	if err := c.Find(nil).Sort("updatedAt").All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddSite inserts an Site in the DB
func (r Repository) AddSite(site Site) bool {

	site.ID = bson.NewObjectId()
	site.UUID = uuid.New().String()
	site.CreatedAt = time.Now()
	site.UpdatedAt = time.Now()
	err := r.s.DB(DBNAME).C(CNAME).Insert(site)

	fmt.Printf("%v", site)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// PatchSite updates an Site in the DB
func (r Repository) PatchSite(match map[string]interface{}, update map[string]interface{}) bool {

	err := r.s.DB(DBNAME).C(CNAME).Update(match, update)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteSite deletes an Site
func (r Repository) DeleteSite(id string) string {

	// Remove user
	if err := r.s.DB(DBNAME).C(CNAME).Remove(bson.M{"uuid": id}); err != nil {
		log.Fatal(err)
		return "500"
	}

	// Write status
	return "200"
}
