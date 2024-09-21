CREATE TYPE role AS ENUM ('Management','DBA', 'Analytics');

CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    roles role[] NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT IdPk PRIMARY KEY (id)
);


-- create index
CREATE INDEX idx_users_name_id ON Users(id, name);