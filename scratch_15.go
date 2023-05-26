package main

import (
	"fmt"
	"log"
	"os"
)

type EnvGetter interface {
	GetEnv(key string) string
}

type LoginFeature struct {
	envGetter EnvGetter
}

func (feat LoginFeature) Login(name, password string) {
	url := feat.envGetter.GetEnv("DOMAIN")
	if url == "" {
		log.Println("Domain is empty")
		return
	}

	fmt.Println("Hitting API:", url)
}

type DummyEnvGetter struct {}

func (d DummyEnvGetter) GetEnv(key string) string {
	return "123"
}

type OsEnvGetter struct {}

func (o OsEnvGetter) GetEnv(key string) string {
	return os.Getenv(key)
}

func main() {
	f := LoginFeature{
		envGetter: DummyEnvGetter{},
	}
	f.Login("iqbal", "password")
	f.envGetter = OsEnvGetter{}
	f.Login("iqbal", "password")
}
