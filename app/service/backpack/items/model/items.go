package model

import (
	"backend/pkg/json"
	"gorm.io/datatypes"
)

const (
	EmptyAttrJSON = "{}"
)

type Item struct {
	ID          int32          `json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Type        int32          `gorm:"not null" json:"type"`
	Price       int32          `gorm:"not null" json:"price"`
	Discount    float32        `gorm:"not null;default:1.0" json:"discount"`
	Available   bool           `gorm:"not null;default:false" json:"available"`
	Attributes  datatypes.JSON `json:"attributes"`
}

func (Item) TableName() string {
	return "np_backpack_items"
}

func (item *Item) SetAttributes(v *Attributes) (err error) {
	item.Attributes, err = json.Marshal(v)
	return
}

func (item *Item) GetAttributes() (v *Attributes, err error) {
	v = &Attributes{}
	if len(item.Attributes) == 0 {
		item.Attributes = []byte(EmptyAttrJSON)
	}
	err = json.Unmarshal(item.Attributes, v)
	return
}
