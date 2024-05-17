package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	DATABASE_USER    string 
	DATABASE_PASSWORD    string 
	DATABASE_NAME    string 
	DATABASE_HOST    string 
	PORT     string 

}

func InitAppConfig() error {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }

	return validateEmptyEnv()
}

func GetEnvs() EnvVars{
	env := EnvVars{}
	env.DATABASE_USER = os.Getenv("DATABASE_USER")
	env.DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	env.DATABASE_NAME = os.Getenv("DATABASE_NAME")
	env.DATABASE_HOST = os.Getenv("DATABASE_HOST")
	env.PORT = os.Getenv("PORT")
	return env

}

func validateEmptyEnv() error {
	required := []string{"PORT", "DATABASE_PASSWORD", "DATABASE_NAME", "DATABASE_HOST", "DATABASE_USER"}

	for _, v := range required {
		if value := os.Getenv(v); value == "" {
			return fmt.Errorf("error empty environment %v", v)
		}
	}
	return nil
}
