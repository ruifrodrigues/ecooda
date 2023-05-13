CREATE TABLE IF NOT EXISTS category_challenge
(
    category_id  INTEGER,
    challenge_id INTEGER,

    CONSTRAINT fk_category_challenge_category_id FOREIGN KEY (category_id) REFERENCES categories (id),
    CONSTRAINT fk_category_challenge_challenge_id FOREIGN KEY (challenge_id) REFERENCES challenges (id),

    PRIMARY KEY(category_id, challenge_id)
);