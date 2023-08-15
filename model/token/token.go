package token

import (
	"github.com/go-yaml/yaml"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type tokenConfig struct {
	Secret string
}

func getConfig(configName string) *tokenConfig {
	config, err := os.ReadFile("config/token.yaml")
	if err != nil {
		log.Fatalln("[FILE READER]", err)
	}

	data := make(map[string]tokenConfig)

	err = yaml.Unmarshal(config, &data)
	if err != nil {
		log.Fatalln("[YAML]", err)
	}

	configData, ok := data[configName]
	if ok == false {
		return nil
	}

	return &configData
}

func New(datas any, config string) (*string, error) {
	auth := jwt.New(jwt.SigningMethodHS256)
	claims := auth.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["values"] = datas

	authStr, err := auth.SignedString([]byte(getConfig(config).Secret))
	if err != nil {
		return nil, err
	}

	return &authStr, nil
}

func IsValid(auth string, config string) (*jwt.Token, error) {

	t, err := jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrInvalidKey
		}

		return []byte(getConfig(config).Secret), nil
	})

	if err != nil {
		return nil, jwt.ErrSignatureInvalid
	}

	if !t.Valid {
		return nil, jwt.ErrTokenSignatureInvalid
	}

	return t, nil
}
