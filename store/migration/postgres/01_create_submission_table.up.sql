CREATE TABLE IF NOT EXISTS submission (
    id CHAR(6) PRIMARY KEY,
    language_id INT NOT NULL,
    status VARCHAR(10) DEFAULT 'pending',
    stdout TEXT DEFAULT '', 
    stderr TEXT DEFAULT '',
    time VARCHAR(10) DEFAULT '0',
    memory VARCHAR(10) DEFAULT '0'
);