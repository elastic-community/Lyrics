package infrastructure

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetData(url string) ([]byte, int, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 404, err
	}
	//req.Header.Set("authorization", "Basic "+auth)

	//resp, err := client.Do(req)ctx
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error Connect to GraphqlServer", err)

	}

	b, err := ioutil.ReadAll(resp.Body)
	return b, resp.StatusCode, err

}
