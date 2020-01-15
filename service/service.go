package service

import (
	"context"
	"github.com/lhlyu/libra/logger"
	"github.com/sirupsen/logrus"
)

type Service struct {
	ctx context.Context
}

func (s Service) Log() *logrus.Entry {
	return logger.GetLogger(s.ctx)
}

func (s Service) Context() context.Context {
	return s.ctx
}
