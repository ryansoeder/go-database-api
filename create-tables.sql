DROP TABLE IF EXISTS albums;
CREATE TABLE albums (
  id         INTEGER PRIMARY KEY ASC,
  title      TEXT,
  artist     TEXT,
  price      INTEGER
);

INSERT INTO albums
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);