package shell_test

import (
	"golang-devcontainer/src/config"
	configMock "golang-devcontainer/src/config/mock"
	"golang-devcontainer/src/service"
	serviceMock "golang-devcontainer/src/service/mock"
	"golang-devcontainer/src/shell"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCalculateShellImpl_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		calculateService service.CalculateService
		cfg              config.Config
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "GetMaxPrintCountとCalculateService.Serviceがそれぞれ一度だけ呼ばれることを確認する",
			fields: fields{
				calculateService: func() service.CalculateService {
					mockCalculate := serviceMock.NewMockCalculateService(ctrl)
					mockCalculate.EXPECT().
						Service(gomock.Any()).
						Return(0).
						Times(1)
					return mockCalculate
				}(),
				cfg: func() config.Config {
					mockConfig := configMock.NewMockConfig(ctrl)
					mockConfig.EXPECT().
						GetMaxPrintCount().
						Return(0).
						Times(1)
					return mockConfig
				}(),
			},
		},
		{
			name: "GetMaxPrintCountから100を受け取った時にCalculateService.Serviceに100が渡ることを確認する",
			fields: fields{
				calculateService: func() service.CalculateService {
					mockCalculate := serviceMock.NewMockCalculateService(ctrl)
					mockCalculate.EXPECT().
						Service(100).
						Return(5050)
					return mockCalculate
				}(),
				cfg: func() config.Config {
					mockConfig := configMock.NewMockConfig(ctrl)
					mockConfig.EXPECT().
						GetMaxPrintCount().
						Return(100)
					return mockConfig
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &shell.CalculateShellImpl{
				CalculateService: tt.fields.calculateService,
				Cfg:              tt.fields.cfg,
			}
			c.Execute()
		})
	}
}
