

-- Table for US States and Canadian Provinces

DROP TABLE IF EXISTS public.state;

CREATE TABLE public.state (
	state_code char(2) NOT NULL,
	state_name varchar(30) NOT NULL,
	PRIMARY KEY (state_code)
);

-- US States

INSERT INTO public.state(state_code, state_name) VALUES( 'AK', 'Alaska');
INSERT INTO public.state(state_code, state_name) VALUES( 'AL', 'Alabama');
INSERT INTO public.state(state_code, state_name) VALUES( 'AR', 'Arkansas');
INSERT INTO public.state(state_code, state_name) VALUES( 'AZ', 'Arizona');
INSERT INTO public.state(state_code, state_name) VALUES( 'CA', 'California');
INSERT INTO public.state(state_code, state_name) VALUES( 'CO', 'Colorado');
INSERT INTO public.state(state_code, state_name) VALUES( 'CT', 'Connecticut');
INSERT INTO public.state(state_code, state_name) VALUES( 'DC', 'District of Columbia');
INSERT INTO public.state(state_code, state_name) VALUES( 'DE', 'Delaware');
INSERT INTO public.state(state_code, state_name) VALUES( 'FL', 'Florida');
INSERT INTO public.state(state_code, state_name) VALUES( 'GA', 'Georgia');
INSERT INTO public.state(state_code, state_name) VALUES( 'HI', 'Hawaii');
INSERT INTO public.state(state_code, state_name) VALUES( 'IA', 'Iowa');
INSERT INTO public.state(state_code, state_name) VALUES( 'ID', 'Idaho');
INSERT INTO public.state(state_code, state_name) VALUES( 'IL', 'Illinois');
INSERT INTO public.state(state_code, state_name) VALUES( 'IN', 'Indiana');
INSERT INTO public.state(state_code, state_name) VALUES( 'KS', 'Kansas');
INSERT INTO public.state(state_code, state_name) VALUES( 'KY', 'Kentucky');
INSERT INTO public.state(state_code, state_name) VALUES( 'LA', 'Louisiana');
INSERT INTO public.state(state_code, state_name) VALUES( 'MA', 'Massachusetts');
INSERT INTO public.state(state_code, state_name) VALUES( 'MD', 'Maryland');
INSERT INTO public.state(state_code, state_name) VALUES( 'ME', 'Maine');
INSERT INTO public.state(state_code, state_name) VALUES( 'MI', 'Michigan');
INSERT INTO public.state(state_code, state_name) VALUES( 'MN', 'Minnesota');
INSERT INTO public.state(state_code, state_name) VALUES( 'MO', 'Missouri');
INSERT INTO public.state(state_code, state_name) VALUES( 'MS', 'Mississippi');
INSERT INTO public.state(state_code, state_name) VALUES( 'MT', 'Montana');
INSERT INTO public.state(state_code, state_name) VALUES( 'NC', 'North Carolina');
INSERT INTO public.state(state_code, state_name) VALUES( 'ND', 'North Dakota');
INSERT INTO public.state(state_code, state_name) VALUES( 'NE', 'Nebraska');
INSERT INTO public.state(state_code, state_name) VALUES( 'NH', 'New Hampshire');
INSERT INTO public.state(state_code, state_name) VALUES( 'NJ', 'New Jersey');
INSERT INTO public.state(state_code, state_name) VALUES( 'NM', 'New Mexico');
INSERT INTO public.state(state_code, state_name) VALUES( 'NV', 'Nevada');
INSERT INTO public.state(state_code, state_name) VALUES( 'NY', 'New York');
INSERT INTO public.state(state_code, state_name) VALUES( 'OH', 'Ohio');
INSERT INTO public.state(state_code, state_name) VALUES( 'OK', 'Oklahoma');
INSERT INTO public.state(state_code, state_name) VALUES( 'OR', 'Oregon');
INSERT INTO public.state(state_code, state_name) VALUES( 'PA', 'Pennsylvania');
INSERT INTO public.state(state_code, state_name) VALUES( 'RI', 'Rhode Island');
INSERT INTO public.state(state_code, state_name) VALUES( 'SC', 'South Carolina');
INSERT INTO public.state(state_code, state_name) VALUES( 'SD', 'South Dakota');
INSERT INTO public.state(state_code, state_name) VALUES( 'TN', 'Tennessee');
INSERT INTO public.state(state_code, state_name) VALUES( 'TX', 'Texas');
INSERT INTO public.state(state_code, state_name) VALUES( 'UT', 'Utah');
INSERT INTO public.state(state_code, state_name) VALUES( 'VA', 'Virginia');
INSERT INTO public.state(state_code, state_name) VALUES( 'VT', 'Vermont');
INSERT INTO public.state(state_code, state_name) VALUES( 'WA', 'Washington');
INSERT INTO public.state(state_code, state_name) VALUES( 'WI', 'Wisconsin');
INSERT INTO public.state(state_code, state_name) VALUES( 'WV', 'West Virginia');
INSERT INTO public.state(state_code, state_name) VALUES( 'WY', 'Wyoming');

INSERT INTO public.state(state_code, state_name) VALUES( 'AP', 'Armed Forces Pacific');
INSERT INTO public.state(state_code, state_name) VALUES( 'AA', 'Armed Forces Americas');
INSERT INTO public.state(state_code, state_name) VALUES( 'AE', 'Armed Forces Europe');

INSERT INTO public.state(state_code, state_name) VALUES( 'AS', 'American Samoa');
INSERT INTO public.state(state_code, state_name) VALUES( 'GU', 'Guam');
INSERT INTO public.state(state_code, state_name) VALUES( 'MP', 'Northern Mariana Islands');
INSERT INTO public.state(state_code, state_name) VALUES( 'PR', 'Puerto Rico');
INSERT INTO public.state(state_code, state_name) VALUES( 'VI', 'Virgin Islands'); 
INSERT INTO public.state(state_code, state_name) VALUES( 'FM', 'Federated States of Micronesia');
INSERT INTO public.state(state_code, state_name) VALUES( 'MH', 'Marshall Islands');
INSERT INTO public.state(state_code, state_name) VALUES( 'PW', 'Palau');

-- Canada

INSERT INTO public.state(state_code, state_name) VALUES( 'AB', 'Alberta');
INSERT INTO public.state(state_code, state_name) VALUES( 'BC', 'British Columbia');
INSERT INTO public.state(state_code, state_name) VALUES( 'MB', 'Manitoba');
INSERT INTO public.state(state_code, state_name) VALUES( 'NB', 'New Brunswick');
INSERT INTO public.state(state_code, state_name) VALUES( 'NL', 'Newfoundland and Labrador');
INSERT INTO public.state(state_code, state_name) VALUES( 'NS', 'Nova Scotia');
INSERT INTO public.state(state_code, state_name) VALUES( 'NT', 'Northwest Territories');
INSERT INTO public.state(state_code, state_name) VALUES( 'NU', 'Nunavut');
INSERT INTO public.state(state_code, state_name) VALUES( 'ON', 'Ontario');
INSERT INTO public.state(state_code, state_name) VALUES( 'PE', 'Prince Edward Island');
INSERT INTO public.state(state_code, state_name) VALUES( 'QC', 'Quebec');
INSERT INTO public.state(state_code, state_name) VALUES( 'SK', 'Saskatchewan');
INSERT INTO public.state(state_code, state_name) VALUES( 'YT', 'Yukon');

SELECT * FROM public.state;
