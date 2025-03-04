CREATE TABLE playlist (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    is_public BOOLEAN DEFAULT FALSE
);
