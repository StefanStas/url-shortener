package api

import (
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
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
            "error": "Internal error",
        })

        return
    }

    urlData := data.(StoreUrlParams)

    url := &models.Url{Url: urlData.Url}
    url, err := a.app.SaveUrl(url)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
            "error": err.Error(),
        })
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
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
            "error": err.Error(),
        })
        return
    }

    if url == nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H {
            "error": "Url not found",
        })
        return
    }

    c.Redirect(http.StatusMovedPermanently, url.Url)
}
