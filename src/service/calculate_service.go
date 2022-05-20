package service

type CalculateService interface {
	Service(loopNumber int) int
}

func NewCalculateService() CalculateService {
	return &CalculateServiceImpl{}
}

type CalculateServiceImpl struct{}

func (c *CalculateServiceImpl) Service(loopNumber int) int {
	result := 0
	for i := 1; i <= loopNumber; i++ {
		result += i
	}
	return result
}
