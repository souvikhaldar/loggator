package logs

import "time"

type Repository interface {
	StoreLog(log string) error
	FetchLog() ([]string, error)
}
type LogData struct {
	TenantId    int       `json:"tenant_id"`
	LogId       int       `json:"log_id"`
	Log         string    `json:"log"`
	CreatedAt   time.Time `json:"created_at"`
	Date        time.Time `json:"date"`
	Time        time.Time `json:"time"`
	LogLevel    string    `json:"log_level"`
	ServiceName string    `json:"service_name"`
	FileName    string    `json:"file_name"`
	PackageName string    `json:"package_name"`
}

// TODO: Add parser for each field
