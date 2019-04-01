package site

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Site represents a music site
type Site struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	URL       string        `bson:"url" json:"url"`
	CreatedAt time.Time     `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updatedAt"`
	IsHealthy bool          `bson:"is_healthy" json:"isHealthy"`
}

//Sites is an array of Site
type Sites []Site
