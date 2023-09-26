CREATE TABLE
    IF NOT EXISTS users (
        id serial,
        name varchar(255),
        age integer,
        PRIMARY KEY (id)
    );
