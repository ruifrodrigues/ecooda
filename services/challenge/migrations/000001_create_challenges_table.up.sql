CREATE TABLE IF NOT EXISTS challenges
(
    id          INTEGER      NOT NULL PRIMARY KEY UNIQUE AUTO_INCREMENT,
    uuid        VARCHAR(32)  NOT NULL UNIQUE,
    name        VARCHAR(128) NOT NULL UNIQUE,
    description VARCHAR(256) NULL,
    street      VARCHAR(128) NULL,
    postcode    VARCHAR(10)  NULL,
    latitude    FLOAT        NULL,
    longitude   FLOAT        NULL,
    thumbnail   VARCHAR(256) NULL,
    gallery     JSON         NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP    NULL
);
