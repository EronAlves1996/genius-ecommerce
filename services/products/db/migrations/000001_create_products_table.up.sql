CREATE TABLE product (
  id serial PRIMARY KEY,
  name varchar(256) NOT NULL,
  price BIGINT NOT NULL,
  image_url varchar(256)
);
