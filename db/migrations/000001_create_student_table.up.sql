CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    nrc text NOT NULL UNIQUE,
    age integer NOT NULL,
    gender text NOT NULL,
    password text NOT NULL
)
