CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE player (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    age INT,
    country VARCHAR(50)
);

CREATE TABLE court (
    id UUID DEFAULT uuid_generate_v4(),
    location VARCHAR(100),
    surface VARCHAR(50),
    PRIMARY KEY (id)
);

CREATE TABLE referee (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    country VARCHAR(50)
);

CREATE TABLE match (
    id SERIAL PRIMARY KEY,
    player1_id INT,
    player2_id INT,
    court_id UUID,
    referee_id INT,
    date TIMESTAMP,
    FOREIGN KEY (player1_id) REFERENCES player(id),
    FOREIGN KEY (player2_id) REFERENCES player(id),
    FOREIGN KEY (court_id) REFERENCES court(id),
    FOREIGN KEY (referee_id) REFERENCES referee(id)
);

CREATE TABLE score (
    match_id INT,
    player_id INT,
    set1 INT,
    set2 INT,
    set3 INT,
    PRIMARY KEY (match_id, player_id),
    FOREIGN KEY (match_id) REFERENCES match(id),
    FOREIGN KEY (player_id) REFERENCES player(id)
);