package models

import (
	"fli-gateway-api/lib"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type LineOAModel struct {}

type LineOA struct {
	bongo.DocumentBase `bson:",inline"`
	Name string `bson:"name" json:"name"`
	ClientID string `bson:"clientID" json:"clientID"`
	ClientSecret string `bson:"clientSecret" json:"clientSecret"`
}

func (LineOAModel) List() (*[]LineOA, error) {
	connection := lib.DBConnection()

	lineOA := &[]LineOA{}

	result := connection.Collection("LINE_OA").Find(bson.M{})
	err := result.Query.All(lineOA)

	return lineOA, err
}
