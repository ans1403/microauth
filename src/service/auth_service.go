package service

import (
	"context"
	"microauth/src/constants"
	"microauth/src/domain"
	"microauth/src/logging"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type AuthService interface {
	SignUp(req *domain.SignUpRequest) error
	ConfirmSignUp(req *domain.ConfirmSignUpRequest) error
	ForgotPassword(req *domain.ForgotPasswordRequest) error
	ConfirmForgotPassword(req *domain.ConfirmForgotPasswordRequest) error
	SignIn(req *domain.SignInRequest) (*cognitoidentityprovider.InitiateAuthOutput, error)
}

func NewAuthService() AuthService {
	app := constants.NewApp()
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(app.AwsDefaultRegion))

	return &authService{
		client:   cognitoidentityprovider.NewFromConfig(cfg),
		clientId: &app.CognitoClientId,
		logger:   logging.NewLogger(),
	}
}

type authService struct {
	client   *cognitoidentityprovider.Client
	clientId *string
	logger   logging.Logger
}

func (s *authService) SignUp(req *domain.SignUpRequest) error {
	_, err := s.client.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
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
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *authService) ConfirmSignUp(req *domain.ConfirmSignUpRequest) error {
	_, err := s.client.ConfirmSignUp(context.TODO(), &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         s.clientId,
		Username:         &req.Username,
		ConfirmationCode: &req.ConfirmationCode,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *authService) ForgotPassword(req *domain.ForgotPasswordRequest) error {
	_, err := s.client.ForgotPassword(context.TODO(), &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: s.clientId,
		Username: &req.Username,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *authService) ConfirmForgotPassword(req *domain.ConfirmForgotPasswordRequest) error {
	_, err := s.client.ConfirmForgotPassword(context.TODO(), &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         s.clientId,
		Username:         &req.Username,
		ConfirmationCode: &req.ConfirmationCode,
		Password:         &req.NewPassword,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s *authService) SignIn(req *domain.SignInRequest) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	res, err := s.client.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: s.clientId,
		AuthParameters: map[string]string{
			"USERNAME": req.Username,
			"PASSWORD": req.Password,
		},
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return res, nil
}
