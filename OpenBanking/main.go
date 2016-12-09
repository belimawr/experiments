package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {

	log.Println("Listening on 0.0.0.0:5000")

	http.HandleFunc("/callback/", callback)
	http.HandleFunc("/", login)

	if err := http.ListenAndServe("0.0.0.0:5000", nil); err != nil {
		log.Fatal("Could not start the server")
	}

}

func callback(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")
	authCode := r.URL.Query().Get("auth_code")

	body := struct {
		UID          string `json:"uid"`
		AuthCode     string `json:"auth_code"`
		DeveloperKey string `json:"developer_key"`
		SecretKey    string `json:"secret_key"`
	}{
		UID:          uid,
		AuthCode:     authCode,
		DeveloperKey: os.Getenv("DEV_KEY"),
		SecretKey:    os.Getenv("SEC_KEY"),
	}

	b, err := json.Marshal(body)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(
		"https://sb-autenticacao-api.original.com.br/OriginalConnect/AccessTokenController",
		"application/json",
		bytes.NewReader(b),
	)

	if err != nil {
		log.Fatal(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	dump, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(dump))

	respMap := map[string]string{}

	err = json.NewDecoder(resp.Body).Decode(&respMap)

	if err != nil {
		log.Fatal(err)
	}

	token := strings.Split(respMap["access_token"], " ")[1]
	accountBalance := balance(token)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(accountBalance))
}

func login(w http.ResponseWriter, r *http.Request) {
	devKey := os.Getenv("DEV_KEY")
	url := fmt.Sprintf("https://sb-autenticacao-api.original.com.br/OriginalConnect?scopes=account&callback_url=http://%s:5000/callback/&callback_id=1&developer_key=%s", getLocalIP(), devKey)

	http.Redirect(w, r, url, 301)
}

func balance(token string) []byte {
	url := "https://sandbox.original.com.br/accounts/v1/balance"
	req, _ := http.NewRequest("get", url, nil)

	req.Header.Add(
		"Authorization",
		token)
	req.Header.Add(
		"developer-key",
		os.Getenv("DEV_KEY"))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	dump, _ := httputil.DumpResponse(res, true)

	fmt.Println(string(dump))
	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return b
}

// Source: http://stackoverflow.com/a/31551220/1535214
// getLocalIP returns the non loopback local IP of the host
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
