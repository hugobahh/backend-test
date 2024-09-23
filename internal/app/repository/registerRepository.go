package repository

import (
	"backend-test/internal/config"
	mysql "backend-test/pkg/database/mysql"
	"backend-test/pkg/logger"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type IRepository interface {
	RegisterEntrance(context.Context, string) error
	RegisterExist(context.Context, string) error
}

type RegRepository struct {
	config         *config.Configuration
	Log            logger.Logger
	mysqlConnector *mysql.MySQLConnector
	statements     map[string]*sql.Stmt
	timeout        time.Duration
}

func NewRegRepository(config *config.Configuration, log *logger.Log) *RegRepository {
	sqlConfig := mysql.DataSource{
		Host:         config.MySQLDataSource.Host,
		User:         config.MySQLDataSource.User,
		Password:     config.MySQLDataSource.Password,
		Name:         config.MySQLDataSource.Name,
		Port:         config.MySQLDataSource.Port,
		ReadOnly:     config.MySQLDataSource.ReadOnly,
		Timeout:      config.MySQLDataSource.Timeout,
		TimeoutQuery: config.MySQLDataSource.TimeoutQuery,
	}
	return &RegRepository{
		Log:            log,
		mysqlConnector: mysql.GetMySQLConnector(&sqlConfig, log),
		timeout:        config.MySQLDataSource.Timeout,
	}
}

func (dr *RegRepository) RegisterEntrance(ctx context.Context, sId string) error {
	//errChk := errors.New("Is here ?")
	//timeout := dr.config.MySQLDataSource.Timeout
	ctx, cancel := context.WithTimeout(ctx, dr.timeout)
	defer cancel()
	//dbConn := dr.mysqlConnector.DBClient
	stmt, ok := dr.mysqlConnector.GetStatements("InsertRegEntrance")
	if !ok {
		return fmt.Errorf("statement 'InsertRegEntrance' not found")
	}

	//_, err := stmt.QueryContext(ctx, sql.Named("Id", sId))
	_, err := stmt.QueryContext(ctx)
	if err != nil {
		logger.GetLogger().Error("RegisterEntrance", "Error executing query", err)
		return err
	}
	return nil

	//return nil
	dr.Log.Infof("Repository-RegisterEntrance...")
	return nil
}

func (dr *RegRepository) RegisterExit(ctx context.Context, sId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dr.config.MySQLDataSource.TimeoutQuery)*time.Second)
	defer cancel()

	rows, err := dr.statements["QueryRegisterEntrance"].QueryContext(ctx,
		sql.Named("Id", sId),
	)
	if err != nil {
		logger.GetLogger().Error("UsuarioRepository", "ObtenerUsuarios", err)
		return err
	}
	defer rows.Close()

	/*	var usuarios []Usuario
		for rows.Next() {
			var usuario Usuario
			if err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Edad); err != nil {
				dr.log.Error("UsuarioRepository", "ObtenerUsuarios", err)
				return nil, err
			}
			usuarios = append(usuarios, usuario)
		}
	*/
	return nil

}
