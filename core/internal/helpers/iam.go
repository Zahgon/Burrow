package helpers

import (
	sarama "github.com/IBM/sarama"
	"github.com/aws/aws-msk-iam-sasl-signer-go/signer"
)

var (
	signerGenerateAuthToken            = signer.GenerateAuthToken
	signerGenerateAuthTokenFromRole    = signer.GenerateAuthTokenFromRole
	signerGenerateAuthTokenFromProfile = signer.GenerateAuthTokenFromProfile
)

type iamTokenProvider struct {
	region, roleArn, profile string
}

func (p *iamTokenProvider) Token() (*sarama.AccessToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
