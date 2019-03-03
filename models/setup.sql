DROP TABLE person;

CREATE TABLE person(
    id serial PRIMARY KEY,
    uuid VARCHAR(255),
    name VARCHAR(255),
    age INTEGER(11),
    created_at timestamp not null
);