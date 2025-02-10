CREATE TABLE IF NOT EXISTS submission (
    id CHAR(6) PRIMARY KEY,
    language_id INT NOT NULL,
    status VARCHAR(10) DEFAULT 'pending',
    stdout TEXT DEFAULT '', 
    stderr TEXT DEFAULT '',
    time DECIMAL(10, 5) DEFAULT 0,
    memory DECIMAL(10, 2) DEFAULT 0
);