package cli

import (
	"fmt"
	"net/http"
)

func commandHello() error {
	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", "http://localhost:8080/test", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
