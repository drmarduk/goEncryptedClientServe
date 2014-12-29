// protocol
package main

import "golang.org/x/crypto/otr"

type Command struct {
	Command   string        `json:"command"`
	User      string        `json:"user,omitempty"`
	Channel   string        `json:"channel,omitempty"`
	Message   string        `json:"message,omitempty"`
	PublicKey otr.PublicKey `json:"publickey,omitempty"`
}
