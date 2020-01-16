package service

import (
	"context"
	"github.com/lhlyu/libra/logger"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Ctx context.Context
}

func (s Service) Log() *logrus.Entry {
	return logger.GetLogger(s.Ctx)
}
