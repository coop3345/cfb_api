package conn

import (
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func APICall[T any](endpoint string, result *T) error {
	token := util.CONFIG.CONNECTIONS.API_TOKEN
	url := util.CONFIG.CONNECTIONS.API_URL_BASE + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(b))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("unmarshaling: %w", err)
	}
	return nil
}
