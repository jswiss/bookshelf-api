CREATE TABLE IF NOT EXISTS books(
  id SERIAL UNIQUE,
  title VARCHAR(100) NOT NULL,
  author VARCHAR(100) NOT NULL,
  cover_image VARCHAR,
  in_stock BOOLEAN default TRUE,
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP default NOW()
);

CREATE TABLE IF NOT EXISTS friends(
  id SERIAL UNIQUE,
  name VARCHAR(100) NOT NULL,
  phone INT,
  email VARCHAR UNIQUE,
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP default NOW()
);

CREATE TABLE IF NOT EXISTS borrowed_books(
  id SERIAL UNIQUE,
  book INT,
  friend INT,
  borrowed_date TIMESTAMP default NOW(),
  returned_date TIMESTAMP,
  created_at TIMESTAMP default NOW(),
  updated_at TIMESTAMP default NOW(),
  CONSTRAINT fk_book
      FOREIGN KEY(book)
	  REFERENCES books(id),
  CONSTRAINT fk_friend
      FOREIGN KEY(friend)
	  REFERENCES friends(id)
);

CREATE INDEX ON "books" ("title");
CREATE INDEX ON "books" ("author");
CREATE INDEX ON "friends" ("name");

-- Add trigger to update in_stock when returned

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON books
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON friends
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON borrowed_books
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
