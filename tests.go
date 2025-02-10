package telegramAPI

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewBotAPI(t *testing.T) {
	token := "test_token"
	bot := NewBot(token)

	if bot.Token != token {
		t.Errorf("Expected token %s, got %s", token, bot.Token)
	}

	expectedURL := "https://api.telegram.org/bot" + token + "/"
	if bot.URL != expectedURL {
		t.Errorf("Expected URL %s, got %s", expectedURL, bot.URL)
	}
}

func TestGetUpdates(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/getUpdates" {
			t.Errorf("Expected path /getUpdates, got %s", r.URL.Path)
		}

		response := `{
			"ok": true,
			"result": [{
				"update_id": 123,
				"message": {
					"message_id": 456,
					"text": "test message",
					"chat": {
						"id": "789"
					}
				}
			}]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create bot with test server URL
	bot := &BotAPI{
		Token: "test_token",
		URL:   server.URL + "/",
	}

	updates, err := bot.GetUpdates(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(updates) != 1 {
		t.Errorf("Expected 1 update, got %d", len(updates))
	}

	update := updates[0]
	if update.UpdateID != 123 {
		t.Errorf("Expected update_id 123, got %d", update.UpdateID)
	}
}

func TestSendMessage(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/sendMessage" {
			t.Errorf("Expected path /sendMessage, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create bot with test server URL
	bot := &BotAPI{
		Token: "test_token",
		URL:   server.URL + "/",
	}

	success, err := bot.SendMessange("123", "test message")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !success {
		t.Error("Expected success to be true")
	}
}
