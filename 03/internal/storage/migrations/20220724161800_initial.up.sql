CREATE TABLE IF NOT EXISTS users (
    id        INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username  VARCHAR(25),
    firstname VARCHAR(50),
    lastname  VARCHAR(50),
    email     VARCHAR(320),
    phone     VARCHAR(15)
)