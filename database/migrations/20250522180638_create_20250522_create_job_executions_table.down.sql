CREATE TABLE job_executions (
    id SERIAL PRIMARY KEY,
    job_id INTEGER REFERENCES jobs(id),
    status VARCHAR(20),
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    duration INTEGER,
    exit_code INTEGER,
    output TEXT,
    error TEXT,
    retry_count INTEGER DEFAULT 0
);
