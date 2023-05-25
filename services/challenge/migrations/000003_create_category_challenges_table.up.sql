CREATE TABLE category_challenges
(
    category_id  INT                                 NOT NULL,
    challenge_id INT                                 NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP                           NULL,

    PRIMARY KEY (category_id, challenge_id),
    CONSTRAINT category_challenges_ibfk_1 FOREIGN KEY (category_id) REFERENCES categories (id),
    CONSTRAINT category_challenges_ibfk_2 FOREIGN KEY (challenge_id) REFERENCES challenges (id)
);

CREATE INDEX challenge_id ON category_challenges (challenge_id);

