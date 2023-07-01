package model

// Reply for the /stats endpoint
// Built from the most frequent request received on the server
// and the number of time this request was shot
type StatsReply struct {
	Request FizzbuzzRequest `json:"request"`
	Count   int             `json:"count"`
}
