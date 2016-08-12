package models

import "encoding/json"

type Book struct {
	Id            string `json:"id" bson:"id"`                         //Book Id, not mongo _id
	TitleText     string `json:"title_text" bson:"title_text"`         //Tittle (Harry Potter)
	PublisherName string `json:"publisher_name" bson:"publisher_name"` //Publisher (Oxford University)
	ImprintName   string `json:"imprint_name" bson:"imprint_name"`     //Imprint (Oxford University)
	Languages     []string `json:"languages" bson:"languages"`         //List of ISO languages ([spa, eng])
	Thematics     []string  `json:"thematics" bson:"thematics"`        //List of thematic ids
	Cover         Image `json:"cover" bson:"cover"`                    //Cover image
}

type Thematic struct {
	Id    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Image Image `json:"image" bson:"image"`
}

type Image struct {
	URL    string `json:"url" bson:"url"`
	Width  int16 `json:"width" bson:"width"`
	Height int16 `json:"height" bson:"height"`
}

func ToJsonString(m interface{}) string {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err) //Should never happen
	}
	return string(bytes)
}

func ToPrettyJsonString(m interface{}) string {
	bytes, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		panic(err) //Should never happen
	}
	return string(bytes)
}