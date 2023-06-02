package utils

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateThreadId() string {
	return uuid.New().String()
}

func HashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// GetEnvironmentVariable is used to get environment variable value
func GetEnvironmentVariable(envName string) (value string, IsPresent bool) {
	value, IsPresent = os.LookupEnv(envName)
	return value, IsPresent
}

// IsFileExist wil return true if file exist
func IsFileExist(filePathName string) bool {
	if _, err := os.Stat(filePathName); err == nil {
		return true
	} else {
		return false
	}
}
