CREATE TABLE account (

  user_id serial PRIMARY KEY,

  username VARCHAR(50) not null,
  email VARCHAR(75) UNIQUE not null,
  password VARCHAR(100) not null,

  create_on TIMESTAMP NOT NULL,
  last_login TIMESTAMP

);