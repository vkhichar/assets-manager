CREATE TABLE assets_maintainance(

	asset_id INT,
	start_date DATE NOT NULL,
	end_date DATE,
	cost NUMERIC(5,2) NOT NULL,

	FOREIGN KEY (asset_id)
	REFERENCES assets(id)
	ON DELETE NO ACTION
	ON UPDATE NO ACTION

);

