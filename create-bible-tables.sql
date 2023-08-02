-- Table to store the topics
DROP TABLE IF EXISTS topics;
CREATE TABLE topics (
  topic_id INTEGER PRIMARY KEY ASC NOT NULL,
  topic_name TEXT NOT NULL
);
-- Inserting sample data into the 'topics' table
INSERT INTO topics (topic_name) VALUES
  ('Abortion'),
  ('Capital punishment (Death penalty)'),
  ('LGBTQ+ rights'),
  ('Animal rights and ethics'),
  ('Immigration and refugees'),
  ('Climate change and environmental ethics'),
  ('Religious freedom versus discrimination');

-- Table to store the Bible verses associated with each topic
DROP TABLE IF EXISTS bible_verses;
CREATE TABLE bible_verses (
  verse_id INTEGER PRIMARY KEY,
  topic_id INTEGER NOT NULL,
  verse_text TEXT NOT NULL,
  supports BOOLEAN NOT NULL, -- Indicates if the verse supports or opposes the topic
  FOREIGN KEY (topic_id) REFERENCES topics(topic_id)
);
-- Inserting sample data into the 'bible_verses' table
INSERT INTO bible_verses (topic_id, verse_text, supports) VALUES
  (1, 'Psalm 139:13-16', true),
  (2, 'Genesis 9:6', true),
  (2, 'Matthew 5:38-39', false),
  (3, 'Genesis 2:24', true),
  (3, 'Romans 1:26-27', false),
  (4, 'Proverbs 12:10', true),
  (5, 'Leviticus 19:33-34', true),
  (5, 'Matthew 25:35', true),
  (6, 'Genesis 1:26-30', true),
  (7, 'Matthew 22:21', true),
  (7, 'Galatians 3:28', true);
