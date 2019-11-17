CREATE TABLE IF NOT EXISTS parties (
	id serial NOT NULL,
	host_id int4 NULL,
	"name" VARCHAR (355) NULL,
	"comments" VARCHAR (500) NULL,
	CONSTRAINT parties_pkey1 PRIMARY KEY (id)
);
