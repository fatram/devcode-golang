package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	cfg               configuration
	PublicKey         *[]byte
	PrivateKey        *[]byte
	PublicKeyReadOne  sync.Once
	PrivateKeyReadOne sync.Once
	cfgOnce           sync.Once
	envFile           *string
)

type configuration struct {
	Title          string        `env-default:"Golang Test Fatur Rahman"`
	AccessTokenTTL time.Duration `env:"ACCESS_TOKEN_TTL" env-default:"1440m" env-upd`
	Port           int           `env:"PORT" env-default:"7007"`
	SecretBytes    string        `env:"SECRET_BYTES" env-default:"secret"`
	DatabaseURI    string        `env:"DATABASE_URI" env-upd env-required`
	PublicKey      string        `env:"PUBLIC_KEY" env-required`
	PrivateKey     string        `env:"PRIVATE_KEY" env-required`
}

func Configuration() configuration {
	if envFile == nil {
		log.Panic(`configuration file is not set. Call ReadConfig("path_to_file") first`)
	}
	err := cleanenv.UpdateEnv(&cfg)
	if err != nil {
		log.Fatalf("Config error %s", err.Error())
	}
	return cfg
}

func (c configuration) GetPublicKey() []byte {
	PublicKeyReadOne.Do(func() {
		signKey, err := os.ReadFile(c.PublicKey)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		PublicKey = &signKey
	})
	return *PublicKey
}

func (c configuration) GetPrivateKey() []byte {
	PrivateKeyReadOne.Do(func() {
		signKey, err := os.ReadFile(c.PrivateKey)
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
		PrivateKey = &signKey
	})
	return *PrivateKey
}

func ReadConfig(file string) {
	cfgOnce.Do(func() {
		envFile = &file
		log.Printf(`Reading config file: "%s"`, *envFile)
		err := cleanenv.ReadConfig(file, &cfg)
		if err != nil {
			log.Print(err)
			err := cleanenv.ReadEnv(&cfg)
			if err != nil {
				log.Print("ksks")
				log.Fatalf("Config error %s", err.Error())
			}
		}
	})
}
