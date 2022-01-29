
DROP TABLE IF EXISTS public.roster;

CREATE TABLE public.roster (
jersey_number smallint NOT NULL PRIMARY KEY,
first_name varchar(20) NOT NULL,
last_name varchar(20) NOT NULL,
position varchar(2) NOT NULL
);

INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES
(2, 'Matt', 'Ryan', 'QB');

INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES
(7, 'Younghoe', 'Koo', 'K');

INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES
(8, 'Kyle', 'Pitts', 'TE');

INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES
(24, 'AJ', 'Terrell', 'CB');

INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES
(45, 'Deion', 'Jones', 'LB');

INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES
(97, 'Grady', 'Jarrett', 'DL');

SELECT * FROM public.roster;

