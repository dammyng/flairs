package setup

var (
	CreateDatabase  = `CREATE DATABASE IF NOT EXISTS alpha_plus DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;`
	UseAlphaPlus    = `USE alpha_plus`
	SetTimeZone     = "SET time_zone = '+00:00';"
	DropDB          = `DROP DATABASE IF EXISTS alpha_plus;`
	CreateUserTable = `CREATE TABLE IF NOT EXISTS users (
		id varchar(255) NOT NULL UNIQUE,
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
		email varchar(100) NOT NULL UNIQUE,
		email_verified_at timestamp NULL DEFAULT NULL,
		password varbinary(255) DEFAULT NULL,
		pin varbinary(255) DEFAULT NULL,
		is_profile_complete tinyint(1) DEFAULT '0',
		phone_number varchar(255) DEFAULT '',
		phone_verified_at timestamp NULL DEFAULT NULL,
		photo longtext,
		type varchar(100) NOT NULL,
		country_id int(10) UNSIGNED DEFAULT NULL,
		created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

	InsertDemoUser = "INSERT INTO `users` (`id`, `first_name`, `last_name`, `address`, `street`, `city`, `postal_code`, `state`, `country`, `referrer`, `ref_code`, `how_did_u_hear_about_us`, `username`, `last_card_request`, `passport`, `id_card`, `bvn`, `wallet`, `customer`, `gender`, `dob`, `email`, `email_verified_at`, `password`, `pin`, `is_profile_complete`, `phone_number`, `phone_verified_at`, `photo`, `type`, `country_id`, `created_at`, `updated_at`) VALUES ('a65a388b-9c94-46f8-a99a-90c4807ce83b', '', '', '', '', '', '', '', '', '', '1EqSqWxQfE4OSRc', '', '', '', '', '', '', 0, 0, '', '2020-07-06 14:32:52', 'someone@flairs.com', '2020-07-06 14:32:52', 0x2432612431302474467758624b78596831454e4b4255707148426235754f65485746576f54756e4943587355676f67516b705150384c306674726943, NULL, 0, '', '2020-07-06 14:32:52', '', '', 0, '2020-07-06 14:32:52', '2020-07-06 14:32:52');"
	SelectDefaultUser = "SELECT * FROM users WHERE id='a65a388b-9c94-46f8-a99a-90c4807ce83b' OR Email='someone@flairs.com';"
	ClearUserTable = "DELETE FROM users;"
)
