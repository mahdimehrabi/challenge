package infrastructures

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env has environment stored
type Env struct {
	ServerPort string
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

// NewEnv creates a new environment
func NewEnv() *Env {
	env := Env{}
	return &env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env.ServerPort = os.Getenv("ServerPort")
	env.DBUsername = os.Getenv("DBUsername")
	env.DBPassword = os.Getenv("DBPassword")
	env.DBHost = os.Getenv("DBHost")
	env.DBPort = os.Getenv("DBPort")
	env.DBName = os.Getenv("DBName")
}
