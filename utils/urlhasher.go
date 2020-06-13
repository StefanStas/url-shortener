package utils

import "math/rand"

func HashUrl(length int) string {
    var pool = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    hash := make([]rune, length)
    for i := range hash {
        hash[i] = pool[rand.Intn(len(pool))]
    }

    return string(hash)
}
