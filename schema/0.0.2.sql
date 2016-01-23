drop table gophergala_tags;
drop table gophergala_learning_resource_tags;

create table gophergala_types (
  id serial primary key,
  name varchar(200)
);

alter table gophergala_learning_resources add column type_id integer;
