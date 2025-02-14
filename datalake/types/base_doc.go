package types

import "github.com/mansio-gmbh/goapiutils/ptr"

type BaseDoc struct {
	ID  string  `json:"_id"`
	Rev *string `json:"_rev,omitempty"`
}

func (b BaseDoc) GetID() string {
	return b.ID
}

func (b BaseDoc) GetRev() string {
	return ptr.OrDefault(b.Rev)
}
