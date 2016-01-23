create table gophergala_reviews (
  id serial primary key,
  user_id integer,
  learning_resource_id integer,
  finished boolean,
  review_date date,
  cost integer,
  course_length integer,
  course_size integer,
  course_interaction integer,
  usefulness integer,
  description varchar(4000)
);
