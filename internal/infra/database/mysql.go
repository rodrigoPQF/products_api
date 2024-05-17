package database

import (
	"log"
	"os"

	"github.com/rodrigoPQF/products_api/internal/config"
	"github.com/upper/db/v4"
	mysql "github.com/upper/db/v4/adapter/mysql"
)

type MysqlDB struct {
	session db.Session
}

func (s *MysqlDB) Close() error {
	return s.session.Close()
}

func (s *MysqlDB) Tx(exec RepositoryTransaction) error {
	return s.session.Tx(func(tx db.Session) error {
		return exec(&MysqlDB{
			session: tx,
		})
	})
} 

func New() (*MysqlDB, error){

	envs := config.GetEnvs()
	connection := mysql.ConnectionURL{
		User: envs.DATABASE_USER,
		Password: envs.DATABASE_PASSWORD,
		Database: envs.DATABASE_NAME,
		Host: envs.DATABASE_HOST,
	}

	
	sess, err := mysql.Open(connection)

	if err != nil {
		log.Fatal(err)
		upDb, err := os.ReadFile("assets/sql/products-table.up.sql")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := sess.SQL().Exec(string(upDb)); err != nil {
			log.Fatal(err)
		}
	}



	if err := sess.Ping() ; err != nil {
		return nil, err
	}

	return &MysqlDB{
		session: sess,
	}, nil
}