PRAGMA foreign_keys = OFF;

DROP TABLE IF EXISTS articles;
CREATE TABLE articles (
  "index" INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT,
  blogger INTEGER NOT NULL,
  content TEXT,
  ref INTEGER NOT NULL DEFAULT 0,
  add_date DATETIME,
  cate INTEGER,
  Comment INTEGER NOT NULL DEFAULT 0,
  Recommend INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX idx_articles_add_date ON articles(add_date);
CREATE INDEX idx_articles_blogger ON articles(blogger);
CREATE INDEX idx_articles_cate ON articles(cate);

DROP TABLE IF EXISTS blogger;
CREATE TABLE blogger (
  "index" INTEGER PRIMARY KEY AUTOINCREMENT,
  id TEXT NOT NULL,
  nick TEXT,
  DOB DATETIME NOT NULL,
  blogname TEXT NOT NULL,
  blogskin INTEGER NOT NULL,
  email TEXT NOT NULL,
  visitor INTEGER NOT NULL DEFAULT 0,
  intro TEXT,
  blogs INTEGER NOT NULL DEFAULT 0,
  IP TEXT,
  TS INTEGER NOT NULL DEFAULT 0,
  Last_log DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  Last_post DATETIME,
  Activate INTEGER NOT NULL DEFAULT 0,
  Reveal INTEGER NOT NULL DEFAULT 1,
  userpic INTEGER NOT NULL DEFAULT 20,
  lang TEXT NOT NULL DEFAULT 'zh_cn',
  widget TEXT
);

DROP TABLE IF EXISTS comment;
CREATE TABLE comment (
  "index" INTEGER PRIMARY KEY AUTOINCREMENT,
  blogger INTEGER,
  article INTEGER NOT NULL,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  author TEXT,
  homepage TEXT,
  add_date DATETIME NOT NULL,
  ip TEXT,
  utype INTEGER NOT NULL DEFAULT 0,
  uid TEXT
);
CREATE INDEX idx_comment_article ON comment(article);
CREATE INDEX idx_comment_blogger ON comment(blogger);

DROP TABLE IF EXISTS links;
CREATE TABLE links (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  blogger INTEGER NOT NULL,
  link TEXT NOT NULL,
  URL TEXT NOT NULL,
  reveal INTEGER NOT NULL DEFAULT 1
);

DROP TABLE IF EXISTS userdefinecategory;
CREATE TABLE userdefinecategory (
  "index" INTEGER PRIMARY KEY AUTOINCREMENT,
  blogger INTEGER NOT NULL,
  cate TEXT NOT NULL,
  files INTEGER NOT NULL DEFAULT 0
);
