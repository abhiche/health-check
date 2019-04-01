package site

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Site represents a music site
type Site struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	UUID      string        `bson:"uuid,omitempty" json:"uuid,omitempty"`
	URL       string        `bson:"url,omitempty" json:"url,omitempty"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	IsHealthy bool          `bson:"is_healthy,omitempty" json:"isHealthy,omitempty"`
}

//Sites is an array of Site
type Sites []Site
