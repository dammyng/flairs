package setup

var (
	CreateDatabase  = `CREATE DATABASE IF NOT EXISTS alpha_transaction DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;`
	UseAlphaWallet    = `USE alpha_transaction`
	SetTimeZone     = "SET time_zone = '+00:00';"
	DropDB          = `DROP DATABASE IF EXISTS alpha_transaction;`
	SQLMode           = "SET SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO'"
	ClearWalletTable    = "DELETE FROM transactions;"

	CreateWalletTable = `CREATE TABLE IF NOT EXISTS transactions (
		id int(10) UNSIGNED NOT NULL,
		created_at timestamp NULL DEFAULT NULL,
		updated_at timestamp NULL DEFAULT NULL,
		deleted_at timestamp NULL DEFAULT NULL,
		account varchar(100) NOT NULL,
		memo varchar(100)  NULL,
		wallet_id varchar(100)  NULL,
		payment_type varchar(100) NOT NULL,
		customer_id varchar(100) NOT NULL,
		tex_ref varchar(100) NOT NULL,
		flw_ref varchar(100) NOT NULL,
		currency varchar(100) NOT NULL,
		status varchar(100) NOT NULL,
		message varchar(100)  NULL,
		type varchar(100) NOT NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;
	  `
)
