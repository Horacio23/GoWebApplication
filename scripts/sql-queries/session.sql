CREATE TABLE session (
    id serial NOT NULL,
    session_id character varying(255),
    member_id integer,
    PRIMARY KEY (id)
);
