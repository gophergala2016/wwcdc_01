create table gophergala_meta (
  version varchar(20)
);

insert into gophergala_meta values('0.0.0');

create table gophergala_users (
  id          bigserial   primary key,
  email       varchar(200)  not null,
  passhash    varchar(200)  not null
);
