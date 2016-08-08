package core

import (
	"github.com/minivac/lebroc/app/models"
	"gopkg.in/mgo.v2"
	"io"
	_ "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/bson"
)

type BookRepository struct {
	ds BookDataSource
}

var _ io.Closer = &BookRepository{} //Ensure closable

func (r *BookRepository ) FindBooks(ids...string) ([]models.Book, error) {
	return r.ds.FindBooks(ids...)
}

func (r *BookRepository ) FindBook(id string) (models.Book, error) {
	return r.ds.FindBook(id)
}

func (r *BookRepository) Close() error {
	return r.ds.Close()
}

type BookDataSource interface {
	FindBooks(ids...string) ([]models.Book, error)
	FindBook(id string) (models.Book, error)
	Close() error
}

type MongoDataSource struct {
	Session *mgo.Session
	db      *mgo.Database
}

var _ BookDataSource = &MongoDataSource{}

func NewMongoDataSource(url string) *MongoDataSource {
	mds := &MongoDataSource{}
	session, err := mgo.Dial(url)
	if (err != nil) {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	mds.Session = session
	mds.db = session.DB("ebooks")
	return mds
}

func (ds *MongoDataSource) FindBooks(ids...string) ([]models.Book, error) {
	query := ds.db.C("books").Find(bson.M{"id" : bson.M{"$in" : ids }, })
	var books []models.Book
	err := query.All(&books)
	return books, err
}

func (ds *MongoDataSource ) FindBook(id string) (models.Book, error) {
	query := ds.db.C("books").Find(bson.M{"id" : id})
	var book models.Book
	err := query.One(&book)
	return book, err
}

func (r *MongoDataSource) Close() error {
	r.Session.Close()
	return nil
}
