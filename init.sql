CREATE TABLE clients (
	id serial NOT NULL,
	first_name character varying(255),
	last_name character varying(255),
	address character varying(255),
	city character varying(255),
	state character varying(255),
	zip character varying(255),
	phone character varying(255),
	email character varying(255),
	entrance_date date,
	last_transaction character varying(255),
	transaction_date date,
	payment numeric,
	notes text,
	PRIMARY KEY (id)
);

CREATE TABLE member
(
  id serial NOT NULL,
  username character varying(255),
  password character varying(255),
  first_name character varying(255),
  PRIMARY KEY (id)
);

CREATE TABLE session (
    id serial NOT NULL,
    session_id character varying(255),
    member_id integer,
    PRIMARY KEY (id)
);

CREATE TABLE transactions (
	id serial NOT NULL,
	client_id character varying(255),
	transaction character varying(255),
	amount character varying(255),
	date date,
	PRIMARY KEY (id)
);
