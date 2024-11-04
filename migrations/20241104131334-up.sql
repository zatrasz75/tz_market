
-- +migrate Up
CREATE TABLE IF NOT EXISTS building (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    year_built INT NOT NULL,
    floors INT
    );


-- +migrate Down
DROP TABLE IF EXISTS building;