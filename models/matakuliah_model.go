package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Matakuliah struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nama_matkul    string             `json:"Nama_matkul,omitempty" bson:"lokasi,omitempty"`
	SKS            string             `json:"SKS,omitempty" bson:"kelurahan,omitempty"`
	Dosen_pengampu string             `json:"Dosen_pengampu,omitempty" bson:"kecamatan,omitempty"`
	Email          string             `json:"Email,omitempty" bson:"kota,omitempty"`
}
