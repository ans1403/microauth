package service

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type UserService struct{}

func (s *UserService) getClient() *cognitoidentityprovider.Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_DEFAULT_REGION")))
	client := cognitoidentityprovider.NewFromConfig(cfg)
	return client
}

func (s *UserService) SignUp(username string, password string, email string) *cognitoidentityprovider.SignUpOutput {
	client := s.getClient()
	res, err := client.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("COGNITO_CLIENT_ID")),
		Username: aws.String(username),
		Password: aws.String(password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
		},
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}

func (s *UserService) ConfirmSignUp(username string, confirmationCode string) *cognitoidentityprovider.ConfirmSignUpOutput {
	client := s.getClient()
	res, err := client.ConfirmSignUp(context.TODO(), &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(os.Getenv("COGNITO_CLIENT_ID")),
		Username:         aws.String(username),
		ConfirmationCode: aws.String(confirmationCode),
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}
