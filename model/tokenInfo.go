package model

import (
	"github.com/satori/go.uuid"
)

type TokenInfo struct {
	ClientID string
	UserID   string
	Scope    string
	Token    string
}

func (ti *TokenInfo) GetClientID() (clientID uuid.UUID, err error) {
	clientID, err = uuid.FromString(ti.ClientID)
	return
}

func (ti *TokenInfo) GetUserID() (userID uuid.UUID, err error) {
	userID, err = uuid.FromString(ti.UserID)
	return
}

func (ti *TokenInfo) UserIsNull() bool {
	if ti.UserID == "" {
		return true
	}

	return false
}
