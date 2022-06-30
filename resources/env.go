package resources

import (
	"os"

	"github.com/joho/godotenv"
)

// a local function as a helper to loading .env File with .env file name and key as a parameter and returning value of the key
func Load_Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	val := os.Getenv(key)

	return val
}
