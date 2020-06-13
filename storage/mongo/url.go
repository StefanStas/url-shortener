package mongo

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "time"
    "url-shortener/models"
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

func (u *Url) SaveUrl(url *models.Url) (*models.Url, error) {
    url.CreatedAt = time.Now()
    url.UpdatedAt = time.Now()

    ctx, _ := context.WithTimeout(context.Background(), QueryTimeout)
    res, err := u.store.Database().Collection(UrlsCollection).InsertOne(ctx, url)
    if err != nil {
        return nil, err
    }

    url.Id = res.InsertedID.(primitive.ObjectID)

    return url, nil
}

func (u *Url) FindByUrl(url string) (*models.Url, error) {
    var result *models.Url
    ctx, _ := context.WithTimeout(context.Background(), QueryTimeout)
    filter := bson.M{"url": url}
    err := u.store.Database().Collection(UrlsCollection).FindOne(ctx, filter).Decode(&result)
    switch err {
    case mongo.ErrNoDocuments:
        return nil, nil
    case nil:
        return result, nil
    default:
        return result, err
    }
}

func (u *Url) FindByHash(hash string) (*models.Url, error) {
    var result *models.Url
    ctx, _ := context.WithTimeout(context.Background(), QueryTimeout)
    filter := bson.M{"hash": hash}
    err := u.store.Database().Collection(UrlsCollection).FindOne(ctx, filter).Decode(&result)
    switch err {
    case mongo.ErrNoDocuments:
        return nil, nil
    case nil:
        return result, nil
    default:
        return result, err
    }
}
