package storage

import "url-shortener/models"

type Store interface {
    UrlStore() UrlStore
}

type UrlStore interface {
    SaveUrl(*models.Url) (*models.Url, error)
    FindByUrl(string) (*models.Url, error)
    FindByHash(string) (*models.Url, error)
}
