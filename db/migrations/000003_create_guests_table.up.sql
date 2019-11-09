CREATE TABLE IF NOT EXISTS guests (
	party_refer int4 NULL,
	id serial NOT NULL,
	first_name text NULL,
	last_name text NULL,
	email text NULL,
	is_attending bool NULL,
	CONSTRAINT guests_pkey PRIMARY KEY (id),
	CONSTRAINT guests_party_refer_parties_id_foreign FOREIGN KEY (party_refer) REFERENCES parties(id) ON UPDATE CASCADE ON DELETE CASCADE
);
