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
  verse_id INTEGER PRIMARY KEY ASC NOT NULL,
  verse_reference TEXT NOT NULL,
  verse_text TEXT NOT NULL
);

-- Inserting sample data into the 'bible_verses' table
INSERT INTO bible_verses (verse_reference, verse_text) VALUES
  ('Psalm 139:13-16', 'For you created my inmost being; you knit me together in my mother’s womb. I praise you because I am fearfully and wonderfully made; your works are wonderful, I know that full well. My frame was not hidden from you when I was made in the secret place, when I was woven together in the depths of the earth. Your eyes saw my unformed body; all the days ordained for me were written in your book before one of them came to be.'),
  ('Genesis 9:6', '“Whoever sheds human blood, by humans shall their blood be shed; for in the image of God has God made mankind.'),
  ('Matthew 5:38-39', '“You have heard that it was said, ‘Eye for eye, and tooth for tooth.’ But I tell you, do not resist an evil person. If anyone slaps you on the right cheek, turn to them the other cheek also.'),
  ('Genesis 2:24', 'That is why a man leaves his father and mother and is united to his wife, and they become one flesh.'),
  ('Romans 1:26-27', 'Because of this, God gave them over to shameful lusts. Even their women exchanged natural sexual relations for unnatural ones. In the same way the men also abandoned natural relations with women and were inflamed with lust for one another. Men committed shameful acts with other men, and received in themselves the due penalty for their error.'),
  ('Proverbs 12:10', 'The righteous care for the needs of their animals, but the kindest acts of the wicked are cruel.'),
  ('Leviticus 19:33-34', '“‘When a foreigner resides among you in your land, do not mistreat them. The foreigner residing among you must be treated as your native-born. Love them as yourself, for you were foreigners in Egypt. I am the Lord your God.'),
  ('Matthew 25:35', 'For I was hungry and you gave me something to eat, I was thirsty and you gave me something to drink, I was a stranger and you invited me in,'),
  ('Genesis 1:26-30', 'Then God said, “Let us make mankind in our image, in our likeness, so that they may rule over the fish in the sea and the birds in the sky, over the livestock and all the wild animals,[a] and over all the creatures that move along the ground.” So God created mankind in his own image, in the image of God he created them; male and female he created them. God blessed them and said to them, “Be fruitful and increase in number; fill the earth and subdue it. Rule over the fish in the sea and the birds in the sky and over every living creature that moves on the ground.” Then God said, “I give you every seed-bearing plant on the face of the whole earth and every tree that has fruit with seed in it. They will be yours for food. 30 And to all the beasts of the earth and all the birds in the sky and all the creatures that move along the ground—everything that has the breath of life in it—I give every green plant for food.” And it was so.'),
  ('Matthew 22:21', '“Caesar’s,” they replied. Then he said to them, “So give back to Caesar what is Caesar’s, and to God what is God’s.”'),
  ('Galatians 3:28','There is neither Jew nor Gentile, neither slave nor free, nor is there male and female, for you are all one in Christ Jesus.');

-- Table to map verses to topics and whether the verse supports the topic
DROP TABLE IF EXISTS map;
CREATE TABLE map (
  id INTEGER PRIMARY KEY ASC NOT NULL,
  topic_id INT NOT NULL,
  verse_id INT NOT NULL,
  supports BOOLEAN,
  FOREIGN KEY (topic_id) REFERENCES topics(topic_id),
  FOREIGN KEY (verse_id) REFERENCES bibl_verses(verse_id)

);

-- Inserting sample data into the 'map' table
INSERT INTO map (topic_id, verse_id, supports) VALUES
  (1, 1, true),
  (2, 2, true),
  (2, 3, false),
  (3, 4, true),
  (3, 5, false),
  (4, 6, true),
  (5, 7, true),
  (5, 8, true),
  (6, 9, true),
  (7, 10, true),
  (7, 11, true)
