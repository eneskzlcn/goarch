package service

import logger "<path-to-core>/logger"

type Repository interface {
}

type Service struct {
	logger     logger.Logger
	repository Repository
}

func New(repository Repository, logger logger.Logger) *Service {
	return &Service{repository: repository, logger: logger}
}
