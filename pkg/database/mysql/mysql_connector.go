package mysql

import (
	config "backend-test/internal/config"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql" // Importa el controlador MySQL
)

type MySQLConnector struct {
	DBClient     *sql.DB
	shutdownOnce sync.Once
	timeout      time.Duration
}

var (
	mySQLConnector *MySQLConnector
	once           sync.Once
)

func GetMySQLConnector(datasource *config.MySQLDataSource) *MySQLConnector {
	once.Do(func() {
		mySQLConnector = NewMySQLConnector(*datasource)
	})
	return mySQLConnector
}

func NewMySQLConnector(datasource config.MySQLDataSource) *MySQLConnector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		datasource.User,
		datasource.Password,
		datasource.Host,
		datasource.Port,
		datasource.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("MySQLConnector", "NewMySQLConnector", err)
	}

	db.SetMaxOpenConns(datasource.MaxOpenConnections)
	db.SetMaxIdleConns(datasource.MaxIdleConnections)

	if err = db.Ping(); err != nil {
		log.Fatal("MySQLConnector", "NewMySQLConnector", err)
	}

	return &MySQLConnector{
		DBClient: db,
		timeout:  datasource.Timeout,
	}
}

func (m *MySQLConnector) HealthCheck() error {
	if m.DBClient != nil {
		if err := m.DBClient.Ping(); err != nil {
			return err
		}
	}
	return nil
}

func (m *MySQLConnector) Shutdown() error {
	var err error
	m.shutdownOnce.Do(func() {
		log.Println("Shutting down MySQL client...")
		if err = m.DBClient.Close(); err != nil {
			log.Println("MySQLConnector", "Shutdown", err)
		} else {
			log.Println("MySQL client is now disconnected")
		}
	})
	return err
}
