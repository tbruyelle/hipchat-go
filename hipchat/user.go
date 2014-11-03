package hipchat

// User represents the HipChat user.
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MentionName string `json:"mention_name"`
	Links       Links  `json:"links"`
}
