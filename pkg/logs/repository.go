package logs

type Repository interface {
	StoreLog(log string) error
	FetchLog() ([]string, error)
}
