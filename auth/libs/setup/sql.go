package setup

var (
	CreateDatabase   = `CREATE DATABASE IF NOT EXISTS alpha_plus DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;`
	UseAlphaPlus=`USE alpha_plus`
	DropDB=`DROP DATABASE IF EXISTS alpha_plus;`
	CreateUserTable = `CREATE TABLE IF NOT EXISTS users (
		id varchar(255) NOT NULL DEFAULT '',
		first_name varchar(255) DEFAULT NULL,
		last_name varchar(255) DEFAULT NULL,
		address varchar(255) DEFAULT NULL,
		street varchar(255) DEFAULT NULL,
		city varchar(255) DEFAULT NULL,
		postal_code varchar(255) DEFAULT NULL,
		state varchar(255) DEFAULT NULL,
		country varchar(255) DEFAULT NULL,
		referrer varchar(255) DEFAULT NULL,
		ref_code varchar(255) DEFAULT NULL,
		how_did_u_hear_about_us varchar(255) DEFAULT NULL,
		username varchar(30) DEFAULT NULL,
		last_card_request varchar(255) DEFAULT NULL,
		passport longtext,
		id_card longtext,
		bvn varchar(11) DEFAULT NULL,
		wallet bigint(20) DEFAULT NULL,
		customer bigint(20) DEFAULT NULL,
		gender varchar(20) DEFAULT NULL,
		dob timestamp NULL DEFAULT NULL,
		email varchar(100) NOT NULL,
		email_verified_at timestamp NULL DEFAULT NULL,
		password varbinary(255) DEFAULT NULL,
		pin varbinary(255) DEFAULT NULL,
		is_profile_complete tinyint(1) DEFAULT '0',
		phone_number varchar(255) NOT NULL,
		phone_verified_at timestamp NULL DEFAULT NULL,
		photo longtext,
		type varchar(100) NOT NULL,
		country_id int(10) UNSIGNED DEFAULT NULL,
		created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

	INSERT_DEMO_USERS = ""
)
