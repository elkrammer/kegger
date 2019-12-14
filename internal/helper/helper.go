// helper functions
package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func FindDifference(a, b []int) (diff []int) {
    m := make(map[int]bool)

    for _, item := range b {
        m[item] = true
    }

    for _, item := range a {
        if _, ok := m[item]; !ok {
            diff = append(diff, item)
        }
    }
    return
}
