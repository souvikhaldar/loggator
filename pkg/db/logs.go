package db

type DB struct {
}

func (d *DB) StoreLog(log string) error {
	return nil
}

func (d *DB) FetchLog() ([]string, error) {
	return nil, nil
}
