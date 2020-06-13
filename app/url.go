package app

import (
    "url-shortener/models"
    "url-shortener/utils"
)

func (a *App) SaveUrl(url *models.Url) (*models.Url, error) {
    existingUrl, err := a.store.UrlStore().FindByUrl(url.Url)
    if err != nil {
        return nil, err
    }

    if existingUrl != nil {
        return existingUrl, err
    }
    url.Hash = utils.HashUrl(5)
    url, err = a.store.UrlStore().SaveUrl(url)
    if err != nil {
        return nil, err
    }
    return url, nil
}

func (a *App) FindUrl(hash string) (*models.Url, error) {
    url, err := a.store.UrlStore().FindByHash(hash)
    if err != nil {
        return nil, err
    }
    return url, nil
}
