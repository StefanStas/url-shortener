package mongo

import (
    "go.mongodb.org/mongo-driver/mongo"
    "url-shortener/storage"
)

type Store interface {
    Client() *mongo.Client
    Database() *mongo.Database
    Connect() error
    UrlStore() storage.UrlStore
}
