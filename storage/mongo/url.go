package mongo

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "url-shortener/storage"
)

var (
    UrlsCollection = "urls"
)

type Url struct {
    store Store
}

func NewUrlStore(store Store) storage.UrlStore {
    return &Url{
        store: store,
    }
}

func (u *Url) GetUrls() ([]string, error) {
    data := []string{}
    ctx, _ := context.WithTimeout(context.Background(), QueryTimeout)
    cur, err := u.store.Database().Collection(UrlsCollection).Find(ctx, bson.D{})
    if err != nil {
        return nil, err
    }

    for cur.Next(ctx) {
        var result bson.M
        err := cur.Decode(&result)
        if err != nil {
            return nil, err
        }
        data = append(data, result["url"].(string))
    }

    return data, nil
}
