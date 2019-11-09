CREATE TABLE IF NOT EXISTS parties (
	id serial NOT NULL,
	"name" text NULL,
	invitation_id int4 NULL,
	invitation_sent timestamptz NULL,
	invitation_opened timestamptz NULL,
	is_attending bool NULL,
	"comments" text NULL,
	host_id int4 NULL,
	CONSTRAINT parties_pkey1 PRIMARY KEY (id)
);
