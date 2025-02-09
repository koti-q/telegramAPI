package telegramApi

import (
	"testing"
)

func TestGetUpdates(t *testing.T) {
	// Create bot instance with test token
	bot := NewBotAPI("test_token")

	// Test getting updates
	updates, err := bot.GetUpdates(0)

	// Verify no errors occurred
	if err != nil {
		t.Errorf("GetUpdates failed: %v", err)
	}

	// Verify updates can be retrieved
	if updates == nil {
		t.Error("Expected updates, got nil")
	}
}
