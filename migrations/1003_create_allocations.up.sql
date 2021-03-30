CREATE TABLE allocations(

	user_id INT,
	asset_id INT,
	date_alloc DATE NOT NULL,
	date_dealloc DATE,

	FOREIGN KEY(user_id) 
	REFERENCES users(id)
	ON DELETE NO ACTION
	ON UPDATE NO ACTION,

	FOREIGN KEY(asset_id) 
	REFERENCES assets(id)
	ON DELETE NO ACTION
	ON UPDATE NO ACTION
	
);
