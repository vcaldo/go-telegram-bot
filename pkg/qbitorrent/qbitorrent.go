package qbitorrent

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func auth() string {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, "https://192.168.1.145:8080/api/v2/auth/login", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("admin", "adminadmin")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s\n", string(resBody))
	return string(resBody)
}
