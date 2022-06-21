package logs

import "time"

type Repository interface {
	StoreLog(log LogData) error
	FetchLog() ([]LogData, error)
}
type LogData struct {
	TenantId    int       `json:"tenant_id"`
	LogId       int       `json:"log_id"`
	Log         string    `json:"log"`
	CreatedAt   time.Time `json:"created_at"`
	Date        string    `json:"date"`
	Time        string    `json:"time"`
	LogLevel    string    `json:"log_level"`
	ServiceName string    `json:"service_name"`
	FileName    string    `json:"file_name"`
	PackageName string    `json:"package_name"`
}
