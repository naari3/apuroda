CREATE TABLE users (
  id uuid PRIMARY KEY,
  name VARCHAR(128),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE files (
  id uuid PRIMARY KEY,
  name VARCHAR(128),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);