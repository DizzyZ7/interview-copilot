CREATE TABLE IF NOT EXISTS questions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    difficulty INT NOT NULL,
    tags TEXT[],
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_questions_tags ON questions USING GIN(tags);
CREATE INDEX IF NOT EXISTS idx_questions_difficulty ON questions(difficulty);
