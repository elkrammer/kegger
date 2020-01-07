CREATE TABLE IF NOT EXISTS users (
	id serial NOT NULL,
	"name" text NULL,
	email text NULL,
	"password" text NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS uix_users_email ON public.users USING btree (email);

INSERT INTO public.users
(id, name, email, "password")
VALUES(1, 'admin', 'admin@admin.com', '$2a$14$PAcGOGXqYlJtp4HTK947KuyketW/L15Fytk1dJKM0w0ZPrSax/YHO');
