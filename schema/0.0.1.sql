create table gophergala_learning_resources (
  id      serial   primary key,
  name    varchar(200) not null,
  url     varchar(2000) not null
);

create table gophergala_learning_resource_languages (
  learning_resource_id integer,
  language_id integer
);

create table gophergala_learning_resource_tags (
  learning_resource_id integer,
  tag_id integer
);

create table gophergala_languages (
  id serial primary key,
  name varchar(200) not null
);

create table gophergala_tags (
  id serial primary key,
  name varchar(200) not null
);
