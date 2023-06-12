CREATE TABLE location_challenges
(
    location_id    INT                                 NOT NULL,
    challenge_uuid VARCHAR(36)                         NOT NULL,

    PRIMARY KEY (challenge_uuid),
    CONSTRAINT location_challenges_location_id_fk
        FOREIGN KEY (location_id) REFERENCES locations (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX location_challenges_location_id_index ON location_challenges (location_id);
CREATE INDEX location_challenges_challenge_uuid_index ON location_challenges (challenge_uuid);

