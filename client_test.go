package outlook

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestNewClient(t *testing.T) {
	dummyURL := "http://dummy.outlook.url"
	t.Run("withInvalidURL", func(t *testing.T) {
		client, err := NewClient("", testDummyToken())
		if err == nil {
			t.Errorf("accept invalid url")
		}
		if client != nil {
			t.Errorf("return invalid client: %#v", client)
		}
	})

	t.Run("withInvalidTokenSource", func(t *testing.T) {
		client, err := NewClient(dummyURL, nil)
		if err == nil {
			t.Errorf("accept invalid token source")
		}
		if client != nil {
			t.Errorf("return invalid client: %#v", client)
		}
	})

	t.Run("positiveCase", func(t *testing.T) {
		client, err := NewClient(dummyURL, testDummyToken())
		if err != nil {
			t.Errorf("client construction failed: %v", err)
		}
		if client == nil {
			t.Errorf("return nil client")
		}
	})
}

func testDummyToken() oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: "<< DUMMY TOKEN >>",
		TokenType:   "Bearer",
	})
}
