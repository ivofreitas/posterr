CREATE DATABASE IF NOT EXISTS strider;

CREATE TABLE IF NOT EXISTS strider.users (
    id VARCHAR(40),
    username VARCHAR(14) NOT NULL,
    followers_count INT,
    following_count INT,
    posts_count INT,
    created_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS strider.posts (
    id              VARCHAR(40) NOT NULL,
    content         TEXT NULL,
    parent          VARCHAR(40) NULL,
    created_at      TIMESTAMP   NULL,
    created_by      VARCHAR(40)  NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (created_by) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS strider.follower (
    id              VARCHAR(40) NOT NULL,
    follow          VARCHAR(40) NOT NULL,
    created_at      TIMESTAMP   NULL,
    PRIMARY KEY (id, follow),
    FOREIGN KEY (id) REFERENCES users(id),
    FOREIGN KEY (follow) REFERENCES users(id)
);

INSERT INTO strider.users
VALUES ("dc5b9b53-0bb1-45d0-9eac-f441dcc16d20", "beltrano.silva", 1, 3, 4, NOW());

INSERT INTO strider.users
VALUES ("0525052d-cb43-44bf-9d1b-54b5fd783399", "fulano.santos", 2, 1, 0, NOW());