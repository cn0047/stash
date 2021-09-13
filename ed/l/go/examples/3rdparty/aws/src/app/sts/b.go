package sts

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func AssumeRole(cfg *aws.Config, assumeRoleArn string) (*sts.AssumeRoleOutput, error) {
	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to get new session, error: %v", err)
	}

	svc := sts.New(sess)

	sessionName := "role"
	r, err := svc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         &assumeRoleArn,
		RoleSessionName: &sessionName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to assume role, error: %v", err)
	}

	return r, nil
}
