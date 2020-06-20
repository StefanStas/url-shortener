package api

import (
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "net/http"
    "url-shortener/models"
)

type StoreUrlParams struct {
    Url    string `form:"url" json:"url" binding:"required,url"`
}

func (a *Api) SaveUrlValidator(c *gin.Context) {
    var form StoreUrlParams

    if err := c.ShouldBind(&form); err != nil {
        // TODO: improve validation error message
        // https://github.com/go-playground/validator/blob/master/_examples/translations/main.go
        //for _, err := range err.(validator.ValidationErrors) {
        //    fmt.Printf("Namespace: %s\n", err.Namespace())
        //    fmt.Printf("Field: %s\n", err.Field())
        //    fmt.Printf("StructNamespace: %s\n", err.StructNamespace())
        //    fmt.Printf("StructField: %s\n", err.StructField())
        //    fmt.Printf("Tag: %s\n", err.Tag())
        //    fmt.Printf("ActualTag: %s\n", err.ActualTag())
        //    fmt.Printf("Kind: %s\n", err.Kind())
        //    fmt.Printf("Type: %s\n", err.Type())
        //    fmt.Printf("Value: %s\n", err.Value())
        //    fmt.Printf("Param: %s\n", err.Param())
        //    fmt.Println("--------------------------")
        //}
        c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
        return
    }

    c.Set("data", form)
}

func (a *Api) SaveUrlCtrl(c *gin.Context) {
    data, ok := c.Get("data")
    if !ok {
        logrus.Error("Accessed not existing save url params")
        errorResponseJson(c, errors.New("accessed not existing save url params"))

        return
    }

    urlData := data.(StoreUrlParams)

    url := &models.Url{Url: urlData.Url}
    url, err := a.app.SaveUrl(url)
    if err != nil {
        logrus.Errorf("api.Api.SaveUrlCtrl error: %+v", err)
        errorResponseJson(c, err)
        return
    }

    c.JSON(http.StatusCreated, gin.H{
       "url": url,
    })
}

func (a *Api) RedirectUrlCtrl(c *gin.Context) {
    hash := c.Param("hash")

    url, err := a.app.FindUrl(hash)
    if err != nil {
        logrus.Errorf("api.Api.RedirectUrlCtrl error: %+v", err)
        errorResponseJson(c, err)
        return
    }

    c.Redirect(http.StatusMovedPermanently, url.Url)
}
