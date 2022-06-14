package logmodel

import (
	"encoding/json"
	"log"
	"os"
)

type Logmodel struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func Logger(name string) *log.Logger {
	return log.New(os.Stdout, name, log.LstdFlags)
}

func Logdata(id string, status string, data string, message string) string {
	j, err := json.Marshal(&Logmodel{
		Id:      id,
		Status:  status,
		Data:    data,
		Message: message,
	})
	if err != nil {
		return ""
	}
	return string(j)
}
