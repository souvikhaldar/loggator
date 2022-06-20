CREATE TYPE loglevel AS ENUM ('ERROR','INFO','DEBUG','CRITICAL');
CREATE TABLE if not exists logs (
	tenant_id INT,
	log_id BIGSERIAL PRIMARY KEY, 
	log TEXT NOT NULL,
	created_at TIMESTAMP default now(),
	date DATE NOT NULL,
	time TIME NOT NULL,
	log_level loglevel DEFAULT 'DEBUG', 
	service_name TEXT,
	file_name TEXT,
	package_name TEXT
);
