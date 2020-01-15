package index_service

import "github.com/lhlyu/libra/controller/dto"

type IGener interface {
	CreateFirstText() string
	CreateSecondText() string
	CreateText() []*NameData
}

type G struct {
	IGener
	Dto *dto.GenerDto
}

func (G) CreateFirstText() string {
	return ""
}

func (G) CreateSecondText() string {
	return ""
}

func (G) CreateText() []*NameData {
	return nil
}

func GFactory(param *dto.GenerDto) IGener {
	switch param.Kind {
	default:
		return nil
	}
	return nil
}
