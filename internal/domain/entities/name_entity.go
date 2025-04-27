package entities

import "time"

type NameEntity struct {
	Id string `json:"id"`
	Name string `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
