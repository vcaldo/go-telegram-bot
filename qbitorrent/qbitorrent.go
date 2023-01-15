package qbitorrent

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func Auth() string {
	host := os.Getenv("QB_HOST")
	url := host + "/api/v2/auth/login"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", os.Getenv("QB_USER"))
	_ = writer.WriteField("password", os.Getenv("QB_PASS"))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	req.Header.Add("Referer", host)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}

	// Implement sid error check
	sid := ""
	for _, cookie := range res.Cookies() {
		if cookie.Name == "SID" {
			sid = cookie.Value
		}
	}

	fmt.Println(sid)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	fmt.Println(string(body))
	return sid

}

func GetTorrents() string {
	host := os.Getenv("QB_HOST")
	url := host + "/api/v2/torrents/info"
	method := "GET"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	req.AddCookie(&http.Cookie{
		Name: "SID", Value: Auth(), MaxAge: 60,
	})
	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	fmt.Println(string(body))
	return string(body[:])
	// return "io"
}

func AddTorrent(torrentUrl string) string {
	host := os.Getenv("QB_HOST")
	url := host + "/api/v2/torrent/add"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("urls", torrentUrl)
	_ = writer.WriteField("root_folder", "true")
	err := writer.Close()

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	req.AddCookie(&http.Cookie{
		Name: "SID", Value: Auth(), MaxAge: 60,
	})
	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	fmt.Println(string(body))
	return string(body[:])
}
