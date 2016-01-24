CREATE TABLE pictures
(
  id serial primary key, 
  learning_resource_id integer,
  picture_url varchar(200),
  name varchar(2000)
);
