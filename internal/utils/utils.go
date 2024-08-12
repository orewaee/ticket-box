package utils

import (
	"log"
	"net/http"
)

func WriteBytes(writer http.ResponseWriter, data []byte) {
	if _, err := writer.Write(data); err != nil {
		log.Println(err)
	}
}

func WriteString(writer http.ResponseWriter, data string) {
	WriteBytes(writer, []byte(data))
}
