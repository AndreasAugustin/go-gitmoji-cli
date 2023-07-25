package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetGitmojis() (gitmojis Gitmojis) {
	res, err := http.Get(ConfigInstance.GitmojisUrl)
	if err != nil {
		fmt.Println("error", err)
	}

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &gitmojis)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return
}
