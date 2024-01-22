CREATE TABLE posts (
  id        int             NOT NULL AUTO_INCREMENT PRIMARY KEY,
  post_id   varchar(255)    NOT NULL, 
  title     varchar(255)    NOT NULL,
  content   varchar(255)    NOT NULL,
  category  varchar(255)    NOT NULL
);