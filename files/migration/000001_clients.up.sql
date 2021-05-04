BEGIN ;

CREATE TABLE clients (
	id serial NOT NULL,
	client_id varchar(50) NOT NULL,
	client_secret varchar(100) NOT NULL,
	channel int NOT NULL,
	status int NOT NULL,
	created_at timestamp with time zone not null default current_timestamp,
	created_by int NOT NULL,
	updated_at timestamp with time zone,
	updated_by int NULL,
	unique (client_id),
	CONSTRAINT client_pk PRIMARY KEY (id)
);

-- default client secret : btg85ip8d3b0vd7nc0dg
INSERT INTO clients (client_id,client_secret,channel,status,created_by) VALUES
('1M6i9GcFqCBbYY0087UD','$argon2id$v=19$m=65536,t=1,p=4$E8+9k7/a1uBnCso69wbZGw$9GJfJdo305/fJ1gexm7flvepl+M5EQGgA3rbh2E+V3I',1,1,1)
;

COMMIT ;