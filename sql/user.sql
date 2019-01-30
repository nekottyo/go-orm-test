CREATE
    database IF NOT EXISTS test;
/*
USE test;

CREATE TABLE items
(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE users
(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    users_ibfk_1 INT ,
    name VARCHAR(255),
    age INT,

    FOREIGN KEY (users_ibfk_1)
        REFERENCES items (id),

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO
    test.items (name)
VALUES
    ("item 1");

INSERT INTO
    test.users (name, age, users_ibfk_1)
VALUES
    ("Michaele", 20, 1),
    ("Lesia", 21, NULL),
    ("Ambrose", 22, NULL),
    ("Domenic", 23, NULL),
    ("Bong", 24, NULL),
    ("Versie", 25, NULL),
    ("Angeli", 25, NULL),
    ("Vince", 24, NULL),
    ("Britni", 24, NULL),
    ("Faye", 23, NULL);
*/
