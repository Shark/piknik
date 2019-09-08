package client

import "time"

const (
	// Version - Piknik version
	Version = "0.9.1"
	// DefaultTTL - Time after the clipboard is considered obsolete, in seconds
	DefaultTTL = 7 * 24 * time.Hour
)

// Conf - Shared config
type Conf struct {
	Connect     string
	EncryptSk   []byte
	EncryptSkID []byte
	Psk         []byte
	SignPk      []byte
	SignSk      []byte
	Timeout     time.Duration
	DataTimeout time.Duration
	TTL         time.Duration
}
