package helpers

import (
	"crypto/sha256"
	"crypto/sha512"

	"github.com/xdg/scram"
)

var SHA256 scram.HashGeneratorFcn = sha256.New
var SHA512 scram.HashGeneratorFcn = sha512.New

type XDGSCRAMClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

func (x *XDGSCRAMClient) Begin(userName, password, authzID string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (x *XDGSCRAMClient) Step(challenge string) (response string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (x *XDGSCRAMClient) Done() bool { _ = "STUB: not implemented"; return false }
