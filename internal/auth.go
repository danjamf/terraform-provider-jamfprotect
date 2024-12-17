package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

var authToken string
var providerDomainname string

func Authenticate(domainname string, clientid string, clientpassword string) error {

	authpayload := "{\"client_id\": \"" + clientid + "\", \"password\": \"" + clientpassword + "\"}"

	url := fmt.Sprintf("https://%s/token", domainname)

	req, err := http.NewRequest("POST", url, strings.NewReader(authpayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Create an HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read and print the response
	fmt.Println("Response status:", resp.Status)
	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	//fmt.Println("Response body:", string(body))

	// Create a variable to store the unmarshalled data
	var tokenResp TokenResponse

	errJson := json.Unmarshal(body, &tokenResp)
	if errJson != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	fmt.Println("Access Token:", tokenResp.AccessToken)

	authToken = tokenResp.AccessToken
	providerDomainname = domainname
	return nil
}

func MakeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	// Add Bearer Token for authentication
	req.Header.Set("Authorization", authToken)
	req.Host = providerDomainname     //swap out domain if something specific is provided
	req.URL.Host = providerDomainname //in both the path AND the host field
	fmt.Print(req)
	resp2, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//defer resp2.Body.Close()

	return resp2, nil
}
