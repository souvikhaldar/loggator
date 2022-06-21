package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/souvikhaldar/loggator/pkg/logs"
)

// NOTE: never hardcode the credentials
// read from vault, env vars, etc instead
const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "loggator"
)

type DB struct {
	db *sql.DB
}

func NewDB() *DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &DB{
		db: db,
	}
}
func (d *DB) StoreLog(log logs.LogData) error {
	query := `INSERT INTO logs(tenant_id,log,date,time,log_level,service_name,file_name,package_name) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := d.db.Exec(
		query,
		log.TenantId,
		log.Log,
		log.Date,
		log.Time,
		log.LogLevel,
		log.ServiceName,
		log.FileName,
		log.PackageName,
	)
	return err
}

func (d *DB) FetchLog() ([]logs.LogData, error) {
	logDatas := make([]logs.LogData, 0)
	query := `SELECT tenant_id,log_id,log,created_at,date,time,log_level,service_name,file_name,package_name from logs`
	rows, err := d.db.Query(
		query,
	)
	if err != nil {
		return logDatas, err
	}
	for rows.Next() {
		var ld logs.LogData
		err := rows.Scan(
			&ld.TenantId,
			&ld.LogId,
			&ld.Log,
			&ld.CreatedAt,
			&ld.Date,
			&ld.Time,
			&ld.LogLevel,
			&ld.ServiceName,
			&ld.FileName,
			&ld.PackageName,
		)
		if err != nil {
			log.Println("Error in scanning logData: ", err)
			continue
		}
		logDatas = append(logDatas, ld)
	}
	return logDatas, nil
}
