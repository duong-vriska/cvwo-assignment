CREATE TABLE IF NOT EXISTS "posts" (
  "id" serial PRIMARY KEY,
  "post_id"  text NOT NULL,
  "title"    text NOT NULL,
  "content"  text NOT NULL,
  "category" text NOT NULL
);