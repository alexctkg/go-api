create schema entity;

create table entity.ent_users (
	use_code serial not null,
    use_email varchar(255) not null,
    use_cnpj  int8  null,
    use_razao_social varchar(60)  null,
    use_password varchar(255) not null,
    use_token text null,
	use_type int not null, --0-super user 1-external app
	use_date_ins timestamp not null DEFAULT now(),
    use_date_del timestamp null,	
	
    CONSTRAINT pk_ent_users PRIMARY KEY (use_code),
	CONSTRAINT fk_use_email UNIQUE (use_email)
);

create schema resources;


create table resources.res_product(
	pro_code serial not null,
	pro_costumermid_cpf_cnpj int8 not null,
	pro_costumermid_email varchar(255) not null,
    pro_costumermid_type int null, --0-cpf 1-cnpj
    pro_status int not null default 0, --0-pending; 1-accept; 2-declined
	pro_status_reason text null, 
	pro_date_updt timestamp null, 
	pro_date_ins timestamp not null DEFAULT now(),
    pro_date_del timestamp null,
    use_code_ext int not null, --user external app, solicitation
    use_code_int int null, --superuser code, response

    CONSTRAINT pk_res_produto PRIMARY KEY (pro_code),
    CONSTRAINT fk_use_code1 FOREIGN KEY (use_code_ext) REFERENCES entity.ent_users(use_code),
    CONSTRAINT fk_use_code2 FOREIGN KEY (use_code_int) REFERENCES entity.ent_users(use_code)

);
