package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mahasiswa struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nama_mhs string             `json:"Nama_mhs" bson:"Nama_mhs" validate:"required"`
	NPM      string             `json:"NPM" bson:"No NPM" validate:"required"`
	Jurusan  string             `json:"Jurusan" bson:"Jurusan" validate:"required"`
	Email    string             `json:"Email" bson:"Email" validate:"required"`
}
