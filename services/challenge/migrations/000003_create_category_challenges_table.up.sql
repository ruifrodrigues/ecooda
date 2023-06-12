CREATE TABLE category_challenges
(
    category_id  INT NOT NULL,
    challenge_id INT NOT NULL,

    PRIMARY KEY (category_id, challenge_id),
    CONSTRAINT category_challenges_category_id_fk
        FOREIGN KEY (category_id) REFERENCES categories (id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT category_challenges_challenge_id_fk
        FOREIGN KEY (challenge_id) REFERENCES challenges (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX category_challenges_category_id_index ON category_challenges (category_id);
CREATE INDEX category_challenges_challenge_id_index ON category_challenges (challenge_id);

