CREATE TABLE if not exists logs (
	tenant_id BIGSERIAL,
	log TEXT,
	created_at TIMESTAMP default now()
);
