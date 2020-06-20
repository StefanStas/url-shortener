package app

import (
    "errors"
    "fmt"
    "url-shortener/models"
    "url-shortener/models/errs"
    "url-shortener/utils"
)

func (a *App) SaveUrl(url *models.Url) (*models.Url, error) {
    existingUrl, err := a.store.UrlStore().FindByUrl(url.Url)
    if err != nil && !errors.Is(err, &errs.NotFoundErr{}) {
        return nil, fmt.Errorf("app.url.App.SaveUrl error: %w", err)
    }

    if existingUrl != nil {
        return existingUrl, err
    }
    url.Hash = utils.HashUrl(5)
    url, err = a.store.UrlStore().SaveUrl(url)
    if err != nil {
        return nil, fmt.Errorf("app.url.App.SaveUrl error: %w", err)
    }
    return url, nil
}

func (a *App) FindUrl(hash string) (*models.Url, error) {
    url, err := a.store.UrlStore().FindByHash(hash)
    if err != nil {
        return nil, fmt.Errorf("app.url.App.FindUrl error: %w", err)
    }
    return url, nil
}
