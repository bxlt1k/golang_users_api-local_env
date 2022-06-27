CREATE TABLE users (
    id              int                     NOT NULL PRIMARY KEY AUTO_INCREMENT,
    firstName       varchar(255)            NOT NULL,
    lastName        varchar(255)            NOT NULL,
    password        varchar(255)            NOT NULL,
);