// handling Webhook data

package hipchat

// Response Types

type WebhookLinks struct {
	Links
}

type Webhook struct {
	WebhookLinks WebhookLinks `json:"links"`
	Name         string       `json:"name"`
	Event        string       `json:"event"`
	Pattern      string       `json:"pattern"`
	URL          string       `json:"url"`
	ID           int          `json:"id,omitempty"`
}

type WebhookList struct {
	Webhooks   []Webhook `json:"items"`
	StartIndex int       `json:"startIndex"`
	MaxResults int       `json:"maxResults"`
	Links      PageLinks `json:"links"`
}

// Request Types

type GetAllWebhooksRequest struct {
	MaxResults int `json:"max-results"`
	StartIndex int `json:"start-index"`
}

type CreateWebhookRequest struct {
	Name    string `json:"name"`
	Event   string `json:"event"`
	Pattern string `json:"pattern"`
	URL     string `json:"url"`
}
