CREATE TABLE IF NOT EXISTS topten(
		id int not null auto_increment primary key,
		cand_name varchar(255) not null,
		cid varchar(64) not null unique,
		cycle varchar(64) not null,
		last_updated varchar(255) not null,
		last_updated_ftv_db datetime default CURRENT_TIMESTAMP,

		industry_code0 varchar(255) not null,
		industry_name0 varchar(255) not null,
		indivs0 int not null,
		pacs0 int not null,
		total0 int not null
	
		industry_code1 varchar(255) not null,
		industry_name1 varchar(255) not null,
		indivs1 int not null,
		pacs1 int not null,
		total1 int not null,
		
		industry_code2 varchar(255) not null,
		industry_name2 varchar(255) not null,
		indivs2 int not null,
		pacs2 int not null,
		total2 int not null,
		
		industry_code3 varchar(255) not null,
		industry_name3 varchar(255) not null,
		indivs3 int not null,
		pacs3 int not null,
		total3 int not null,

		industry_code4 varchar(255) not null,
		industry_name4 varchar(255) not null,
		indivs4 int not null,
		pacs4 int not null,
		total4 int not null,

		industry_code5 varchar(255) not null,
		industry_name5 varchar(255) not null,
		indivs5 int not null,
		pacs5 int not null,
		total5 int not null,

		industry_code6 varchar(255) not null,
		industry_name6 varchar(255) not null,
		indivs6 int not null,
		pacs6 int not null,
		total6 int not null,

		industry_code7 varchar(255) not null,
		industry_name7 varchar(255) not null,
		indivs7 int not null,
		pacs7 int not null,
		total7 int not null,

		industry_code8 varchar(255) not null,
		industry_name8 varchar(255) not null,
		indivs8 int not null,
		pacs8 int not null,
		total8 int not null,

		industry_code9 varchar(255) not null,
		industry_name9 varchar(255) not null,
		indivs9 int not null,
		pacs9 int not null,
		total9 int not null);
CREATE UNIQUE INDEX cand_crp on topten (cid);
