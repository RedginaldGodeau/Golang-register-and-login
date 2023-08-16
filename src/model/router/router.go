package router

import (
	"gopkg.in/yaml.v2"
	"log"
	"main/src/domain/handlers"
	"net/http"
	"os"
	"reflect"
)

type route struct {
	Pattern  string   `yaml:"pattern"`
	Function string   `yaml:"function"`
	Methods  []string `yaml:"methods"`
	Redirect string   `yaml:"redirect"`
}

func haveMethod(methods []string, method string) bool {
	accept := false

	for _, met := range methods {
		if method == met {
			accept = true
		}
	}

	return accept
}

func invokeHandler(pattern string, fu string, methods []string, redirect string) {
	var h handlers.Handler

	http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if !haveMethod(methods, req.Method) {
			http.Redirect(w, req, redirect, http.StatusSeeOther)
			return
		}

		args := make([]reflect.Value, 2)
		args[0] = reflect.ValueOf(w)
		args[1] = reflect.ValueOf(req)
		reflect.ValueOf(&h).MethodByName(fu).Call(args)
	})
}

func RouteMaker() {
	dir, _ := os.Getwd()
	config, err := os.ReadFile(dir + "/config/route.yaml")
	if err != nil {
		log.Fatalln("[FILE READER]", err)
	}

	data := make(map[string]route)

	err = yaml.Unmarshal(config, &data)
	if err != nil {
		log.Fatalln("[YAML]", err)
	}

	for _, v := range data {
		pattern := v.Pattern
		function := v.Function
		methods := v.Methods
		redirect := v.Redirect

		if len(methods) == 0 {
			methods = append(methods, "GET")
		}

		invokeHandler(pattern, function, methods, redirect)
	}
}
