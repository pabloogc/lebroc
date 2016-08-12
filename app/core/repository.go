package core

import (
	"github.com/minivac/lebroc/app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
)

type BookRepository struct {
	BookDataSource
}

var _ io.Closer = &BookRepository{} //Ensure closable

func (r *BookRepository) Close() error {
	return r.BookDataSource.Close()
}


//########################################
//MONGODB DATA SOURCE IMPLEMENTATION
//########################################

type BookDataSource interface {
	FindBooks(ids...string) ([]models.Book, error)
	FindBook(id string) (models.Book, error)
	FindBooksWithThematic(thematicId string, offset int, limit int) ([]models.Book, error)

	FindAllThematics() ([]models.Thematic, error)
	Close() error
}

type MongoDataSource struct {
	session *mgo.Session
	db      *mgo.Database
}

//Ensure Mongo is a valid data source
var _ BookDataSource = &MongoDataSource{}

func NewMongoDataSource(url string) *MongoDataSource {
	mds := &MongoDataSource{}
	session, err := mgo.Dial(url)
	if (err != nil) {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	mds.session = session
	mds.db = session.DB("ebooks")
	return mds
}

func (ds *MongoDataSource) FindBooks(ids...string) ([]models.Book, error) {
	query := ds.db.C("books").Find(bson.M{"id" : bson.M{"$in" : ids }, })
	var books []models.Book
	err := query.All(&books)
	return books, err
}

func (ds *MongoDataSource) FindBook(id string) (models.Book, error) {
	query := ds.db.C("books").Find(bson.M{"id" : id})
	var book models.Book
	err := query.One(&book)
	return book, err
}

func (ds *MongoDataSource) FindBooksWithThematic(thematicId string, offset int, limit int) ([]models.Book, error) {
	query := ds.db.C("books").Find(bson.M{"thematics" : thematicId, })
	var books []models.Book
	err := query.Sort("title_text").Skip(offset).Limit(limit).All(&books)
	return books, err
}

func (ds *MongoDataSource) FindAllThematics() ([]models.Thematic, error) {
	query := ds.db.C("thematics").Find(nil)
	var thematics []models.Thematic
	err := query.All(&thematics)
	return thematics, err
}

func (r *MongoDataSource) Close() error {
	r.session.Close()
	return nil
}
