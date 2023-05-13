CREATE TABLE IF NOT EXISTS challenge_location
(
    challenge_id INTEGER,
    location_id  INTEGER,

    CONSTRAINT fk_challenge_location_challenge_id FOREIGN KEY (challenge_id) REFERENCES challenges (id),

    PRIMARY KEY(challenge_id, location_id)
);