package services

import "github.com/hayrullahcansu/cachy/framework/logging"

type BaseService struct {
	logger *logging.Logger
}

func NewBaseService() *BaseService {
	return &BaseService{
		logger: logging.Instance(),
	}
}
