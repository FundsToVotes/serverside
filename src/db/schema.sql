create table if not exists topten (
    id int not null auto_increment primary key,
    cand_name varchar(255) not null,
    cid varchar(64) not null,
    cycle varchar(64) not null,
    last_updated varchar(255) not null,
    last_updated_ftv_db DATETIME not null,
     
    industry_code1 varchar(255) not null,
    industry_name1 varchar(255) not null,
    indivs1 int not null,
    pacs1 int not null,
    total1 int not null,
);
CREATE UNIQUE INDEX cand_name_index
on topten (cand_name);
