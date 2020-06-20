package api

import (
    "errors"
    "github.com/gin-gonic/gin"
    "net/http"
    "url-shortener/models/errs"
)

func getStatusAndMessage(err error) (int, string) {
    switch  {
    case errors.Is(err, &errs.NotFoundErr{}):
        var e *errs.NotFoundErr
        errors.As(err, &e)
        return http.StatusNotFound, e.Error()
    default:
        return http.StatusInternalServerError, "Something went wrong"
    }
}

func errorResponseJson(c *gin.Context, err error)  {
    status, errMsg := getStatusAndMessage(err)
    c.AbortWithStatusJSON(status, gin.H {
        "error": errMsg,
    })
}
