package shippo

import "time"

type ObjectMeta struct {
	ObjectID      string    `json:"object_id"`
	ObjectOwner   string    `json:"object_owner"`
	ObjectState   string    `json:"object_state"`
	ObjectCreated time.Time `json:"object_created"`
	ObjectUpdated time.Time `json:"object_updated"`
}
