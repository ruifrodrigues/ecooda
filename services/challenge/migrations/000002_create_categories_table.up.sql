CREATE TABLE categories
(
    id         INT AUTO_INCREMENT PRIMARY KEY KEY,
    uuid       VARCHAR(36)                         NOT NULL,
    name       VARCHAR(128)                        NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP                           NULL,

    CONSTRAINT id UNIQUE (id),
    CONSTRAINT name UNIQUE (name),
    CONSTRAINT uuid UNIQUE (uuid)
);

