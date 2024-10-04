CREATE TYPE time_period AS ENUM ('جاهلي', 'أموي', 'عباسي', 'أندلسي', 'عثماني ومملوكي', 'متأخر وحديث');

CREATE TABLE IF NOT EXISTS poets (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL,
    bio VARCHAR(500) NOT NULL,
    reviewed BOOLEAN DEFAULT true,
    time_period time_period NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT PoetsIdPk PRIMARY KEY (id)
);


-- create index
CREATE INDEX idx_poets_name_id ON Poets(id, name);