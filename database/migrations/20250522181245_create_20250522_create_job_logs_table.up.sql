CREATE TABLE job_logs (
    id SERIAL PRIMARY KEY,
    job_id INTEGER REFERENCES jobs(id),
    execution_id INTEGER REFERENCES job_executions(id),
    log_time TIMESTAMPTZ,
    log_level VARCHAR(20),
    message TEXT
);

