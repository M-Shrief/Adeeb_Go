
CREATE TABLE IF NOT EXISTS chosen_verses (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    poet_id UUID NOT NULL,
    poem_id UUID NOT NULL,
    verses verse[] NOT NULL,
    reviewed BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (poet_id) REFERENCES poets(id),
    FOREIGN KEY (poem_id) REFERENCES poems(id)
    -- id is unique by constraint
)