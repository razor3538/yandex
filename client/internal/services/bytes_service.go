package services

import (
	"bytes"
	config "client/init"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ByteService struct{}

func NewByteService() *ByteService {
	return &ByteService{}
}

func (bs *ByteService) Save(byte, name, meta string) (int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/byte")

	postBody, _ := json.Marshal(map[string]interface{}{
		"meta":      meta,
		"bytes":     byte,
		"name_pair": name,
	})

	responseBody := bytes.NewReader(postBody)

	req, _ := http.NewRequest("POST", requestURL, responseBody)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", string(token))

	res, err := client.Do(req)

	if err != nil {
		return 500, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}

func (bs *ByteService) Get(name string) (string, int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/get_byte")

	postBody, _ := json.Marshal(map[string]string{
		"name": name,
	})

	responseBody := bytes.NewReader(postBody)

	req, _ := http.NewRequest("POST", requestURL, responseBody)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", string(token))

	res, err := client.Do(req)

	if err != nil {
		return "", 500, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b), res.StatusCode, nil
}

func (bs *ByteService) Delete(name string) (string, int, error) {
	token, err := ioutil.ReadFile("cred.txt")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	requestURL := fmt.Sprintf(config.Env.ApiURL + "/byte")

	postBody, _ := json.Marshal(map[string]string{
		"name": name,
	})

	responseBody := bytes.NewReader(postBody)

	req, _ := http.NewRequest("DELETE", requestURL, responseBody)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	req.Header.Set("Authorization", string(token))

	res, err := client.Do(req)

	if err != nil {
		return "", 500, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b), res.StatusCode, nil
}
