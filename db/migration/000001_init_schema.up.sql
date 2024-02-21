CREATE TABLE IF NOT EXISTS app_user(
  id serial PRIMARY KEY,
  email varchar NOT NULL,
  password varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  first_name varchar NOT NULL,
  last_name varchar NOT NULL
);