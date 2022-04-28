
DROP TABLE IF EXISTS test.test_table1;
DROP TABLE IF EXISTS test.test_table2;
DROP TABLE IF EXISTS test.test_table3;

CREATE TABLE test.test_table1 (
	id int NOT NULL, -- Sequential ID
	display_name varchar(10) NOT NULL,
	sort_order smallint NOT NULL, 
	modified_date timestamp NOT NULL, 
	CONSTRAINT test_table1_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE test.test_table1 IS 'This Is Test Table 1';

-- Column comments

COMMENT ON COLUMN test.test_table1.id IS 'Sequential ID';
COMMENT ON COLUMN test.test_table1.display_name IS 'Name To Display';
COMMENT ON COLUMN test.test_table1.sort_order IS 'The order to display records in';
COMMENT ON COLUMN test.test_table1.modified_date IS 'Date/Time of last update to record';

CREATE TABLE test.test_table2 (
	id int4 NOT NULL,
	amount numeric(5, 2) NOT NULL,
	is_deleted bool NOT NULL,
	CONSTRAINT test_table2_pkey PRIMARY KEY (id)
);

CREATE TABLE test.test_table3 (
	id int4 NOT NULL,
	name varchar(15) NOT NULL,
	expiration_date date NULL,
	state_code char(2) NULL,
	notes text NULL,
	CONSTRAINT test_table3_pkey PRIMARY KEY (id)
);