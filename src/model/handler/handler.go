package handler

import (
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
)

type Handler struct{}

type ressource struct {
	Pattern string `yaml:"pattern"`
	Path    string `yaml:"path"`
}

func InitRessources() {
	dir, _ := os.Getwd()
	config, err := os.ReadFile(dir + "/config/ressources.yaml")
	if err != nil {
		log.Fatalln("[FILE READER]", err)
	}

	data := make(map[string]ressource)

	err = yaml.Unmarshal(config, &data)
	if err != nil {
		log.Fatalln("[YAML]", err)
	}

	for _, v := range data {
		pattern := v.Pattern
		path := v.Path
		http.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(path))))
	}
}
