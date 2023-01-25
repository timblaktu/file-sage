package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type SmugMugConfig struct {
	APIKey             string `required:"true" split_words:"true"`
	APISecret          string `required:"true" split_words:"true"`
	UserToken          string `required:"true" split_words:"true"`
	UserSecret         string `required:"true" split_words:"true"`
	Destination        string `required:"true" split_words:"true"`
	FileNames          string `required:"true" split_words:"true"`
	UseMetadataTimes   bool   `required:"true" split_words:"true"`
	ForceMetadataTimes bool   `required:"true" split_words:"true"`
}

type Config struct {
	Debug   bool
	Timeout time.Duration
	HomeDir string
	Smugmug SmugMugConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error(), "Error loading configuration from .env file")
	}

	var conf Config
	err = envconfig.Process("wupdd", &conf)
	if err != nil {
		log.Fatal(err.Error(), "Error loading Global Config from environment")
	}
	jsonconf, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Config:\n%s\n", jsonconf)
}
