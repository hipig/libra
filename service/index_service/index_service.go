package index_service

import (
	"context"
	"github.com/lhlyu/libra/controller/dto"
	"github.com/lhlyu/libra/response"
	"github.com/lhlyu/libra/service"
	"sort"
)

type indexService struct {
	service.Service
}

func NewIndexService(ctx context.Context) *indexService {
	return &indexService{
		service.Service{
			ctx,
		},
	}
}

func (s *indexService) GenerNames(param *dto.GenerDto) *response.R {
	list := GFactory(param).CreateNames()
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Gender < list[j].Gender
	})
	return response.Success.WithData(list)
}
