DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS comments;


CREATE TABLE posts (
   id serial PRIMARY KEY,
   content TEXT NOT NULL,
   author varchar(255)
);

CREATE TABLE comments (
   commentsId serial PRIMARY KEY, -- primary key column
   content TEXT NOT NULL,
   author VARCHAR(255) NOT NULL,
   post_id INTEGER REFERENCES posts(id)
);