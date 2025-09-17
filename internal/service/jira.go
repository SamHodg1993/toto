package service

import (
	"context"
	"fmt"
	"time"

	"net/http"
)

func StartCallbackServer(expectedState string) (string, error) {
	codeChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		// Make sure the state is what we expect
		if state != expectedState {
			errorChan <- fmt.Errorf("invalid state parameter")
			return
		}

		fmt.Println("We have the users creds. Store in keyring next!")

		// Success, tell the user to close the browser window
		w.WriteHeader(200)
		w.Write([]byte("Authentication successful! You can close this window."))

		// Send code to channel
		codeChan <- code
	})

	// Start server
	server := &http.Server{Addr: ":8989"} // Port is specific to the app. Configured in atlassian developer console
	go server.ListenAndServe()

	// Wait for callback or timeout
	select {
	case code := <-codeChan:
		server.Shutdown(context.Background())
		return code, nil
	case err := <-errorChan:
		server.Shutdown(context.Background())
		return "", err
	case <-time.After(2 * time.Minute):
		server.Shutdown(context.Background())
		return "", fmt.Errorf("authentication timeout")
	}
}
