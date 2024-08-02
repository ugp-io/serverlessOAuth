package ups

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func requestToken() (string, error) {

	// Encode the client ID and client secret using base64
	auth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("UPS_CLIENTID") + ":" + os.Getenv("UPS_SECRETKEY")))

	// Prepare the token request payload
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	// Create a new HTTP request
	req, err := http.NewRequest(
		"POST",
		// "https://onlinetools.ups.com/security/v1/oauth/token",
		"https://wwwcie.ups.com/security/v1/oauth/token",
		strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+auth)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to obtain token, status code: %d", resp.StatusCode)
	}

	// Parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse UPSTokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
