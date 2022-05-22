package service

import (
	"context"
	"microauth/src/constants"
	"microauth/src/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type AuthService struct {
	client   *cognitoidentityprovider.Client
	clientId *string
}

func NewAuthService() *AuthService {
	app := constants.NewApp()
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(app.AwsDefaultRegion))

	s := &AuthService{}
	s.client = cognitoidentityprovider.NewFromConfig(cfg)
	s.clientId = &app.CognitoClientId
	return s
}

func (s *AuthService) SignUp(req *domain.SignUpRequest) *cognitoidentityprovider.SignUpOutput {
	res, err := s.client.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId: s.clientId,
		Username: &req.Username,
		Password: &req.Password,
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: &req.Email,
			},
		},
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}

func (s *AuthService) ConfirmSignUp(req *domain.ConfirmSignUpRequest) *cognitoidentityprovider.ConfirmSignUpOutput {
	res, err := s.client.ConfirmSignUp(context.TODO(), &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         s.clientId,
		Username:         &req.Username,
		ConfirmationCode: &req.ConfirmationCode,
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}

func (s *AuthService) ForgotPassword(req *domain.ForgotPasswordRequest) *cognitoidentityprovider.ForgotPasswordOutput {
	res, err := s.client.ForgotPassword(context.TODO(), &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: s.clientId,
		Username: &req.Username,
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}

func (s *AuthService) ConfirmForgotPassword(req *domain.ConfirmForgotPasswordRequest) *cognitoidentityprovider.ConfirmForgotPasswordOutput {
	res, err := s.client.ConfirmForgotPassword(context.TODO(), &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         s.clientId,
		Username:         &req.Username,
		ConfirmationCode: &req.ConfirmationCode,
		Password:         &req.NewPassword,
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}

func (s *AuthService) SignIn(req *domain.SignInRequest) *cognitoidentityprovider.InitiateAuthOutput {
	res, err := s.client.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: s.clientId,
		AuthParameters: map[string]string{
			"USERNAME": req.Username,
			"PASSWORD": req.Password,
		},
	})

	if err != nil {
		panic(err.Error())
	}

	return res
}
