package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Nilai struct {
	Id           primitive.ObjectID `json:"id,omitempty"`
	NPM_ms       string             `json:"NPM_ms,omitempty" validate:"required"`
	Presensi     string             `json:"Presensi,omitempty" validate:"required"`
	Nilai_akhir  string             `json:"Nilai_akhir,omitempty" validate:"required"`
	Grade        string             `json:"Grade,omitempty" validate:"required"`
	Tahun_ajaran string             `json:"tahun_ajaran,omitempty" validate:"required"`
}
