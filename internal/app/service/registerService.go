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
	repository repository.RegRepository
}

func NewRegService(repo *repository.RegRepository, log *logger.Log) *RegisterService {
	return &RegisterService{Log: log, repository: *repo}
}

func (rs *RegisterService) RegisterEntrance(ctx context.Context, sId string) error {
	//return s.repository.GetItems(ctx)
	//rs.Log.Infof("RegisterEntrance ...")
	err := rs.repository.RegisterEntrance(ctx, "1")
	if err != nil {
		logger.GetLogger().Error("RegisterService", "Error executing query", err)
		return err
	}
	return nil
}

func (rs *RegisterService) RegisterExit(ctx context.Context, sId string) error {
	rs.Log.Infof("RegisterExit ...")
	rs.repository.RegisterExit(ctx, "2")
	return nil
}
