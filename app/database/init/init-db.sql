CREATE TABLE IF NOT EXISTS books(
  id SERIAL UNIQUE NOT NULL,
  title VARCHAR(100) NOT NULL,
  author VARCHAR(100) NOT NULL,
  cover_image VARCHAR default 'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fmattsko.files.wordpress.com%2F2012%2F02%2Fbook-cover1500s.jpg&f=1&nofb=1' NOT NULL,
  in_stock BOOLEAN default TRUE NOT NULL,
  created_at TIMESTAMP default NOW() NOT NULL,
  updated_at TIMESTAMP default NOW() NOT NULL
);

CREATE TABLE IF NOT EXISTS friends(
  id SERIAL UNIQUE NOT NULL,
  full_name VARCHAR(100) NOT NULL,
  photo VARCHAR NOT NULL default '',
  created_at TIMESTAMP default NOW() NOT NULL,
  updated_at TIMESTAMP default NOW() NOT NULL
);

CREATE TABLE IF NOT EXISTS borrowed_books(
  id SERIAL UNIQUE NOT NULL,
  book_id INT NOT NULL,
  friend_id INT NOT NULL,
  borrowed_date TIMESTAMP default NOW() NOT NULL,
  returned_date TIMESTAMP,
  created_at TIMESTAMP default NOW() NOT NULL,
  updated_at TIMESTAMP default NOW() NOT NULL,
  CONSTRAINT fk_book_id
      FOREIGN KEY(book_id)
	  REFERENCES books(id),
  CONSTRAINT fk_friend_id
      FOREIGN KEY(friend_id)
	  REFERENCES friends(id)
);

CREATE INDEX ON "books" ("title");
CREATE INDEX ON "books" ("author");
CREATE INDEX ON "friends" ("full_name");

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

CREATE OR REPLACE FUNCTION out_of_stock()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE books
  SET in_stock = FALSE
  WHERE id = NEW.book_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER change_stock_on_borrowed_insert
AFTER INSERT ON borrowed_books
FOR EACH ROW
EXECUTE PROCEDURE out_of_stock();
