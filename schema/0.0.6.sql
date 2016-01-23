CREATE TABLE pictures
(
  picture_id serial NOT NULL,
  learning_resource_id integer,
  picture_url varchar(200),
  name varchar(2000)
);
