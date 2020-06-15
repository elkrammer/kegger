CREATE TABLE IF NOT EXISTS settings (
    name VARCHAR (255) NULL,
    value VARCHAR (255) NULL,
    description VARCHAR (255) NULL,
    CONSTRAINT settings_pkey PRIMARY KEY (name)
);

INSERT INTO public.settings
("name", value, description)
VALUES('event_name', 'Kegger''s Bash', 'Event Name');

INSERT INTO public.settings
("name", value, description)
VALUES('event_location', '1180 6th Ave, New York, NY 10036, United States', 'Event Location');

INSERT INTO public.settings
("name", value, description)
VALUES('event_date', '2020-11-19 14:00:00.000', 'Event Date');

INSERT INTO public.settings
("name", value, description)
VALUES('dress_code', 'Casual', 'Dress Code');

INSERT INTO public.settings
("name", value, description)
VALUES('groom_name', 'Fat Mike', 'Groom''s Name');

INSERT INTO public.settings
("name", value, description)
VALUES('bride_name', 'Beer', 'Bride''s Name');

INSERT INTO public.settings
("name", value, description)
VALUES('time_zone', 'America/Toronto', 'Event time zone');

-- Invite Settings
INSERT INTO public.settings
("name", value, description)
VALUES('invite_background', '/uploads/default_bg.jpg', 'Invitation Background Image Path');
