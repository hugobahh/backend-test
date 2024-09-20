package repository

import (
	"backend-test/pkg/logger"
	"context"
	"time"
)

type IRepository interface {
	RegisterEntrance(context.Context, string) error
	RegisterExist(context.Context, string) error
}

type DataRepository struct {
	Log     logger.Logger
	timeout time.Duration
}

func NewResumeRepository(log *logger.Log) *DataRepository {
	return &DataRepository{
		Log: log,
		//timeout: config.SQLDataSource.Timeout,
	}
}

func (dr *DataRepository) RegisterEntrance(ctx context.Context, sId string) error {
	//errChk := errors.New("Is here ?")
	dr.Log.Infof("Repository-RegisterEntrance...")
	return nil
}
func (dr *DataRepository) RegisterExit(ctx context.Context, sId string) error {
	//errChk := errors.New("Is here ?")
	dr.Log.Infof("Repository-RegisterExit...")
	return nil
}
