CREATE TABLE IF NOT EXISTS users (
	id serial NOT NULL,
	"name" text NULL,
	email text NULL,
	"password" text NULL,
	host_refer int4 NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX uix_users_email ON public.users USING btree (email);
