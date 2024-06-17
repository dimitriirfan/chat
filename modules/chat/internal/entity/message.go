package entity

type Message struct {
	ID        string      `json:"id"`
	Content   string      `json:"content"`
	Sender    Participant `json:"sender"`
	SessionID string      `json:"session_id"`
}
