package config

import (
	"os"
)

type Configuration struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string
	DB_TIMEZONE string
}

func (c *Configuration) LoadEnvironmentVars() {

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Unable to retrieve environment variables")
	// }

	// fmt.Println(os.Environ())

	c.DB_HOST = os.Getenv("DB_HOST")
	c.DB_USER = os.Getenv("DB_USER")
	c.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	c.DB_NAME = os.Getenv("DB_NAME")
	c.DB_PORT = os.Getenv("DB_PORT")
	c.DB_SSLMODE = os.Getenv("DB_SSLMODE")
	c.DB_TIMEZONE = os.Getenv("DB_TIMEZONE")

}
