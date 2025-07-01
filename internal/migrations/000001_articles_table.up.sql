CREATE TABLE IF NOT EXISTS articles(
    id SERIAL PRIMARY KEY,
    title varchar(255) unique,
    content TEXT
);

CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    name varchar(255) unique
);

INSERT INTO categories (name) VALUES ('games'), ('journey'), ('family');

CREATE TABLE IF NOT EXISTS article_categories(
    article_id int NOT NULL REFERENCES articles (id) ON DELETE CASCADE,
    category_id int NOT NULL REFERENCES categories (id) ON DELETE CASCADE
);