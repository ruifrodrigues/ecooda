CREATE TABLE locations
(
    id         INT AUTO_INCREMENT PRIMARY KEY KEY,
    uuid       VARCHAR(36)                         NOT NULL,
    parent_id  INT       DEFAULT NULL,
    name       VARCHAR(128)                        NOT NULL,
    type       TINYINT                             NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP                           NULL,

    CONSTRAINT id UNIQUE (id),
    CONSTRAINT uuid UNIQUE (uuid),

    CONSTRAINT location_location_parent_id_fk FOREIGN KEY (parent_id) REFERENCES locations (id)
);

