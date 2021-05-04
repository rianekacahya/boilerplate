BEGIN ;

CREATE TABLE referentials (
	id serial NOT NULL,
	ref_id int NOT NULL,
	table_name varchar(50) NOT NULL,
	coloum varchar(50) NOT NULL,
	value varchar(100) NOT NULL,
	created_at timestamp with time zone not null default current_timestamp,
	created_by int NOT NULL,
	updated_at timestamp with time zone,
	updated_by int NULL,
	CONSTRAINT referentials_pk PRIMARY KEY (id)
);

INSERT INTO referentials (ref_id,table_name,coloum,value,created_by) VALUES
(1,'client','channel','WEB',1)
,(2,'client','channel','ANDROID',1)
,(3,'client','channel','IOS',1)
,(1,'client','status','ENABLE',1)
,(2,'client','status','DISABLE',1)
;

COMMIT ;