package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Settings struct {
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ParseSettingsFile(fileName string) (Settings, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return Settings{}, err
	}
	log.Println("[INFO]: reading settings file")
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	var settings Settings
	json.Unmarshal(bytes, &settings)
	return settings, nil
}
