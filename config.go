package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type LocalConfig struct {
	RootPath string `required:"true" split_words:"true" default:""`
}

func (c *LocalConfig) Specified() bool {
	// Parens required around struct initializer to resolve parsing ambiguity
	//   https://go.dev/ref/spec#Composite_literals
	return *c != (LocalConfig{})
}
func (c *LocalConfig) Valid() bool {
	f, err := os.Open(c.RootPath)
	if err != nil {
		log.Fatalf("Can't open LocalConfig.RootPath: %s: %s", c.RootPath, err)
	}
	var finfo os.FileInfo
	finfo, err = f.Stat()
	if err != nil {
		log.Fatalf("Invalid LocalConfig.RootPath %s: %s", c.RootPath, err)
	}
	return finfo.IsDir()
}

type SmugMugConfig struct {
	URL                string `required:"true" split_words:"true" default:""`
	APIKey             string `required:"true" split_words:"true" default:""`
	APISecret          string `required:"true" split_words:"true" default:""`
	UserToken          string `required:"true" split_words:"true" default:""`
	UserSecret         string `required:"true" split_words:"true" default:""`
	Destination        string `required:"true" split_words:"true" default:""`
	FileNames          string `required:"true" split_words:"true" default:""`
	UseMetadataTimes   bool   `required:"true" split_words:"true" default:false`
	ForceMetadataTimes bool   `required:"true" split_words:"true" default:false`
}

func (c *SmugMugConfig) Specified() bool {
	// Parens required around struct initializer to resolve parsing ambiguity
	//   https://go.dev/ref/spec#Composite_literals
	return *c != (SmugMugConfig{})
}

func (c *SmugMugConfig) Valid() bool {
	// TODO: fix validation
	return true
}

type Config struct {
	Debug   bool
	Timeout time.Duration
	HomeDir string
	Local   LocalConfig
	Smugmug SmugMugConfig
}

// The struct returned is a singleton.
func loadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading configuration from .env file: %s", err.Error())
	}

	var conf Config
	err = envconfig.Process("WDD", &conf)
	if err != nil {
		log.Fatalf("Error loading Global Config from environment: %s", err.Error())
	}
	jsonconf, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		log.Fatalf("Error Marshalling conf struct: %s", err.Error())
	}
	log.Printf("Loaded Config:\n%s\n", jsonconf)
	return conf
}
