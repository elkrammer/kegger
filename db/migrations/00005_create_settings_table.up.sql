CREATE TABLE IF NOT EXISTS settings (
    name VARCHAR (255) NULL,
    value VARCHAR (255) NULL,
    description VARCHAR (255) NULL,
    CONSTRAINT settings_pkey PRIMARY KEY (name)
);

INSERT INTO public.settings
("name", value, description)
VALUES('event_name', 'WingDing''s Bash', 'Event Name');

INSERT INTO public.settings
("name", value, description)
VALUES('event_date', '2020-11-19 14:00:00.000', 'Event Date');
