package service

import (
	"context"
	"microauth/src/domain"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type UserService struct {
	client   *cognitoidentityprovider.Client
	clientId *string
}

func NewUserService() *UserService {
	awsDefaultRegion := os.Getenv("AWS_DEFAULT_REGION")
	cognitoClientId := os.Getenv("COGNITO_CLIENT_ID")

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsDefaultRegion))

	s := &UserService{}
	s.client = cognitoidentityprovider.NewFromConfig(cfg)
	s.clientId = &cognitoClientId
	return s
}

func (s *UserService) SignUp(req *domain.SignUpRequest) *cognitoidentityprovider.SignUpOutput {
	res, err := s.client.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId: s.clientId,
		Username: req.Username,
		Password: req.Password,
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: req.Email,
			},
		},
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}

func (s *UserService) ConfirmSignUp(req *domain.ConfirmSignUpRequest) *cognitoidentityprovider.ConfirmSignUpOutput {
	res, err := s.client.ConfirmSignUp(context.TODO(), &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         s.clientId,
		Username:         req.Username,
		ConfirmationCode: req.ConfirmationCode,
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}

func (s *UserService) ForgotPassword(req *domain.ForgotPasswordRequest) *cognitoidentityprovider.ForgotPasswordOutput {
	res, err := s.client.ForgotPassword(context.TODO(), &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: s.clientId,
		Username: req.Username,
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}

func (s *UserService) ConfirmForgotPassword(req *domain.ConfirmForgotPasswordRequest) *cognitoidentityprovider.ConfirmForgotPasswordOutput {
	res, err := s.client.ConfirmForgotPassword(context.TODO(), &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         s.clientId,
		Username:         req.Username,
		ConfirmationCode: req.ConfirmationCode,
		Password:         req.NewPassword,
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}
