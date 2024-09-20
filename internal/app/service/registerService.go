package service

import (
	"backend-test/internal/app/repository"
	"backend-test/pkg/logger"
	"context"
)

type IRegisterService interface {
	//UpdateCandidateResume(context.Context, string) error
	RegisterEntrance(context.Context, string) error
	RegisterExit(context.Context, string) error
}

type RegisterService struct {
	Log        logger.Logger
	repository repository.DataRepository
}

func NewResumeService(repo *repository.DataRepository, log *logger.Log) *RegisterService {
	return &RegisterService{Log: log, repository: *repo}
}

func (rs *RegisterService) RegisterEntrance(ctx context.Context, sId string) error {
	//return s.repository.GetItems(ctx)
	rs.Log.Infof("RegisterEntrance ...")
	rs.repository.RegisterEntrance(ctx, "1")
	return nil
}

func (rs *RegisterService) RegisterExit(ctx context.Context, sId string) error {
	rs.Log.Infof("RegisterExit ...")
	rs.repository.RegisterExit(ctx, "2")
	return nil
}
