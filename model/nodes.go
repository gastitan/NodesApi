package model

import (
	"errors"
)

type NodeType string

const (
	Store  NodeType = "Store"
	Pickup NodeType = "Pickup"
)

type Node struct {
	ID           int      `json:"id" validate:"required,gte=0" example:"1"`
	NodeType     NodeType `json:"nodeType" validate:"required" example:"Store"`
	Location     Location `json:"location" validate:"required"`
	Address      string   `json:"address" validate:"required" example:"Balcarce 50"`
	BusinessHour string   `json:"businessHour" validate:"-" example:"8-18"`
	Capacity     int8     `json:"capacity" validate:"-" example:"100"`
}

type Location struct {
	Lat float64 `json:"lat" validate:"required" example:"34.00"`
	Lng float64 `json:"lng" validate:"required" example:"-25.00"`
}

type Nearest struct {
	Distance float64 `json:"distance"`
	Node     Node    `json:"node"`
}

func (st NodeType) IsValid() error {
	switch st {
	case Store, Pickup:
		return nil
	}

	return errors.New("invalid node type")
}
