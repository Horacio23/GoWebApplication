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
