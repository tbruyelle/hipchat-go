package hipchat

import (
	"net/http"
)

// ClientCredentials represents the OAuth2 client ID and secret for an integration
type ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

// OAuthAccessToken represents a newly created Hipchat OAuth access token
type OAuthAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint32 `json:"expires_in"`
	GroupID     uint32 `json:"group_id"`
	GroupName   string `json:"group_name"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GenerateTokenOptions specifies the optionnal parameters to Client.GenerateToken
// method.
type GenerateTokenOptions struct {
	// The user name to generate a token on behalf of.
	// Only valid in the 'password' and 'client_credentials' grants.
	Username string `url:"username,omitempty"`

	// The type of grant request.
	GrantType string `url:"grant_type,omitempty"`

	// The authorization code to exchange for an access token.
	//	Only valid in the 'authorization_code' grant.
	Code string `url:"code,omitempty"`

	// The name of the public oauth client retrieving a token for.
	// Only valid in the 'authorization_code' and 'refresh_token' grants.
	ClientName string `url:"client_name,omitempty"`

	// The URL that was used to generate an authorization code, and it must match
	// that value. Only valid in the 'authorization_code' grant.
	RedirectURI string `url:"redirect_uri,omitempty"`

	// A space-delimited list of scopes that is requested.
	Scope string `url:"scope,omitempty"`

	// The user's password to use for authentication when creating a token.
	// Only valid in the 'password' grant.
	Password string `url:"password,omitempty"`

	// The name of the group to which the related user belongs.
	// Only valid in the 'authorization_code' and 'refresh_token' grants.
	GroupID string `url:"group_id,omitempty"`

	// The refresh token to use to generate a new access token.
	// Only valid in the 'refresh_token' grant.
	RefreshToken string `url:"refresh_token,omitempty"`
}

// GenerateToken returns back an access token for a given integration's client ID and client secret
//
//  HipChat API documentation: https://www.hipchat.com/docs/apiv2/method/generate_token
func (c *Client) GenerateToken(opt *GenerateTokenOptions) (*OAuthAccessToken, *http.Response, error) {
	req, err := c.NewRequest("POST", "oauth/token", opt, nil)
	if err != nil {
		return nil, nil, err
	}

	token := new(OAuthAccessToken)
	resp, err := c.Do(req, token)
	if err != nil {
		return nil, resp, err
	}
	c.authToken = token.AccessToken
	return token, resp, nil
}

const (
	// ScopeAdminGroup - Perform group administrative tasks
	ScopeAdminGroup = "admin_group"

	// ScopeAdminRoom - Perform room administrative tasks
	ScopeAdminRoom = "admin_room"

	// ScopeImportData - Import users, rooms, and chat history. Only available for select add-ons.
	ScopeImportData = "import_data"

	// ScopeManageRooms - Create, update, and remove rooms
	ScopeManageRooms = "manage_rooms"

	// ScopeSendMessage - Send private one-on-one messages
	ScopeSendMessage = "send_message"

	// ScopeSendNotification - Send room notifications
	ScopeSendNotification = "send_notification"

	// ScopeViewGroup - View users, rooms, and other group information
	ScopeViewGroup = "view_group"

	// ScopeViewMessages - View messages from chat rooms and private chats you have access to
	ScopeViewMessages = "view_messages"

	// ScopeViewRoom - View room information and participants, but not history
	ScopeViewRoom = "view_room"
)
