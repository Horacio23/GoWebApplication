CREATE TABLE member
(
  id serial NOT NULL,
  username character varying(255),
  password character varying(255),
  first_name character varying(255),
  PRIMARY KEY (id)
);