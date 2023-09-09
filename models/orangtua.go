package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrangTua struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	Nama_ortu    string             `json:"Nama_ortu,omitempty" validate:"required"`
	Phone_number string             `json:"Phone_number,omitempty" validate:"required"`
	Email        string             `json:"Email,omitempty" validate:"required"`
}
