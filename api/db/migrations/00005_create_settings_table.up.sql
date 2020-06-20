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
VALUES('wedding_website_url', 'http://127.0.0.1:8081', 'Wedding Website URL');

INSERT INTO public.settings
("name", value, description)
VALUES('kegger_frontend_url', 'http://127.0.0.1:8080', 'Kegger Frontend URL');

INSERT INTO public.settings
("name", value, description)
VALUES('kegger_api_url', 'http://127.0.0.1:4040', 'Kegger API URL');

-- Invite Settings
INSERT INTO public.settings
("name", value, description)
VALUES('invite_image_en', '/uploads/default_bg.jpg', 'English Invitation Image Path');

INSERT INTO public.settings
("name", value, description)
VALUES('invite_image_es', '/uploads/default_bg.jpg', 'Spanish Invitation Image Path');
