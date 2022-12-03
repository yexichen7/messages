package model

type Message struct {
	Mid      int64  `json:"mid"`
	AuthorID string `json:"authorID"`
	RecID    string `json:"recID"`
	Detail   string `json:"detail"`
}
