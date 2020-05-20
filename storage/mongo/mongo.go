package mongo

import (
    "context"
    "fmt"
    "github.com/sirupsen/logrus"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "time"
    "url-shortener/storage"
)

const (
    ConnectTimeout = 5 * time.Second
    QueryTimeout = 5 * time.Second
)

type Mongo struct {
    config   *Config
    client   *mongo.Client
    database *mongo.Database

    urlStore storage.UrlStore
}

type Config struct {
    Host   string
    Port   int
    DbName string
}

func InitMongoStore(config *Config) (storage.Store, error) {
    store := NewMongoStore(config)
    if err := store.Connect(); err != nil {
        return nil, err
    }

    store.urlStore = NewUrlStore(store)

    return store, nil
}

func NewMongoStore(config *Config) *Mongo {
    return &Mongo{
        config: config,
    }
}

func (m *Mongo) Connect() error {
    ctx, _ := context.WithTimeout(context.Background(), ConnectTimeout)

    client, err := mongo.Connect(ctx, options.Client().
        ApplyURI(fmt.Sprintf("mongodb://%s:%d", m.config.Host, m.config.Port)))
    if err != nil {
        return err
    }

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        return err
    }

    m.client = client
    m.database = client.Database(m.config.DbName)

    logrus.Info("Connected to DB")
    return nil
}

func (m *Mongo) Client() *mongo.Client {
    return m.client
}

func (m *Mongo) Database() *mongo.Database {
    return m.database
}

func (m *Mongo) UrlStore() storage.UrlStore {
    return m.urlStore
}
