package contacts

type Contact struct {
	ContactID int64  `json:"contact_id"`
	Slack     string `json:"slack"`
	Instagram string `json:"instagram"`
	X         string `json:"x"`
	Facebook  string `json:"facebook"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
