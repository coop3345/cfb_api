package conn

import (
	"cfbapi/util"
	"fmt"
	"io"
	"net/http"
	"time"
)

func APICall(endpoint string) ([]byte, error) {
	token := util.API_Token
	url := util.API_URL_BASE + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Add the Bearer token to the Authorization header
	req.Header.Add("Authorization", "Bearer "+token)

	// Use a client to send the request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(b))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	return body, err
}
