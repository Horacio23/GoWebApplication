CREATE TABLE transactions (
	id serial NOT NULL,
	client_id character varying(255),
	transaction character varying(255),
	amount character varying(255),
	date date,
	PRIMARY KEY (id)
);
