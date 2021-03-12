CREATE TABLE assets(

	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL, 
	category VARCHAR(50) NOT NULL,
	specification JSON, 
	init_cost NUMERIC(7,2) NOT NULL,
	status INT NOT NULL

);
