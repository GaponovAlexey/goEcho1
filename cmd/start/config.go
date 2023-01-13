package start

import (
	"log"
	"os"

	"github.com/joho/godotenv"

)

type ConfigDatabase struct {
	Port     string
	Host     string
	Name     string
	User     string
	Password string
}

var cfg = &ConfigDatabase{
	Port: os.Getenv("PORT"),
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	log.Println("PORT -> ", cfg.Port)
	log.Println("PORT2 -> ", os.Getenv("PORT"))
}
