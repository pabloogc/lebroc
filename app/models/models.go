package models

import "encoding/json"

type Book struct {
	Id            string `json:"id" bson:"id"`
	TitleText     string `json:"title_text" bson:"title_text"`
	PublisherName string `json:"publisher_name" bson:"publisher_name"`
	ImprintName   string `json:"imprint_name" bson:"imprint_name"`
	Languages     []string `json:"languages" bson:"languages"`
}

func NewBook(id string, title string) *Book {
	b := &Book{}
	b.Id = id
	b.TitleText = title
	b.PublisherName = "Grupo Planeta"
	b.ImprintName = "Roca"
	b.Languages = []string{"es", "eng", "ger"}
	return b
}

func ToJsonString(m interface{}) string {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err) //Should never happen
	}
	return string(bytes)
}