CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  post_id  varchar(255) NOT NULL, 
  title    text NOT NULL,
  content  text NOT NULL,
  category text NOT NULL
);