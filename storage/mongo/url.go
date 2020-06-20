package mongo

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "time"
    "url-shortener/models"
    "url-shortener/models/errs"
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

func (u *Url) SaveUrl(url *models.Url) (*models.Url, error) {
    url.CreatedAt = time.Now()
    url.UpdatedAt = time.Now()

    ctx, _ := context.WithTimeout(context.Background(), QueryTimeout)
    res, err := u.store.Database().Collection(UrlsCollection).InsertOne(ctx, url)
    if err != nil {
        return nil, fmt.Errorf("storage.mongo.url.SaveUrl error: %w", err)
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
        return nil, fmt.Errorf("storage.mongo.url.FindByUrl by value <%s> error: %w", url, &errs.NotFoundErr{ EntityType: "url" })
    case nil:
        return result, nil
    default:
        return result, fmt.Errorf("storage.mongo.url.FindByUrl error: %w", err)
    }
}

func (u *Url) FindByHash(hash string) (*models.Url, error) {
    var result *models.Url
    ctx, _ := context.WithTimeout(context.Background(), QueryTimeout)
    filter := bson.M{"hash": hash}
    err := u.store.Database().Collection(UrlsCollection).FindOne(ctx, filter).Decode(&result)
    switch err {
    case mongo.ErrNoDocuments:
        return nil, fmt.Errorf("storage.mongo.url.FindByHash with hash value: <%s> error: %w", hash, &errs.NotFoundErr{ EntityType: "url" })
    case nil:
        return result, nil
    default:
        return result, fmt.Errorf("storage.mongo.url.FindByHash with hash value: <%s> error: %w", hash, err)
    }
}
