CREATE TYPE verse AS (
    first VARCHAR(50),
    sec VARCHAR(50)
);


CREATE TABLE IF NOT EXISTS poems (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    intro VARCHAR(50) UNIQUE NOT NULL,
    poet_id UUID NOT NULL,
    verses verse[] NOT NULL,
    reviewed BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (poet_id) REFERENCES poets(id),

    CONSTRAINT PoemsIdPk PRIMARY KEY (id)
);

-- create index
CREATE INDEX idx_poems_intro_id ON Poems(id, intro);