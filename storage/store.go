package storage

type Store interface {
    UrlStore() UrlStore
}

type UrlStore interface {
    GetUrls() ([]string, error)
}
