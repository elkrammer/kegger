CREATE TABLE IF NOT EXISTS guests (
	id serial NOT NULL,
	party_refer int4 NULL,
	first_name VARCHAR (60) NULL,
	last_name VARCHAR (60) NULL,
	email VARCHAR (120) NULL,
	is_attending bool NULL,
	invitation_id VARCHAR (27) NULL,
	invitation_sent timestamptz NULL,
	invitation_opened timestamptz NULL,
	CONSTRAINT guests_pkey PRIMARY KEY (id),
	CONSTRAINT guests_party_refer_parties_id_foreign FOREIGN KEY (party_refer) REFERENCES parties(id) ON UPDATE CASCADE ON DELETE CASCADE
);
