CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE priority AS ENUM ('low', 'medium', 'high');

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       due_date TIMESTAMP,
                       priority priority DEFAULT 'medium',
                       done BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE attachments (
                             id SERIAL PRIMARY KEY,
                             task_id INTEGER REFERENCES tasks(id) ON DELETE CASCADE,
                             filename VARCHAR(255),
                             file_path VARCHAR(255),
                             uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE logs (
                      id SERIAL PRIMARY KEY,
                      endpoint VARCHAR(255),
                      method VARCHAR(10),
                      status_code INTEGER,
                      requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
