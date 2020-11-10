package setup

var (
	CreateDatabase  = `CREATE DATABASE IF NOT EXISTS alpha_wallet DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;`
	UseAlphaWallet    = `USE alpha_wallet`
	SetTimeZone     = "SET time_zone = '+00:00';"
	DropDB          = `DROP DATABASE IF EXISTS alpha_wallet;`
	SQLMode           = "SET SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO'"
	ClearWalletTable    = "DELETE FROM wallets;"

	CreateWalletTable = `CREATE TABLE IF NOT EXISTS wallets (
		id varchar(255) NOT NULL UNIQUE,
		user_id varchar(255) NOT NULL,
		account_bal varchar(255) NOT NULL,
		available_bal varchar(255) DEFAULT NULL,
		ledger_bal varchar(255) DEFAULT NULL,
		wallet_type varchar(1) NOT NULL,
		name varchar(20) NOT NULL,
		memo varchar(255) NOT NULL,
		currency varchar(20) NOT NULL,
		status varchar(255) DEFAULT NULL,
		date_created varchar(255) DEFAULT NULL,
		date_bal_update varchar(255) DEFAULT NULL,
		last_update timestamp NULL DEFAULT NULL,
		deleted_at timestamp NULL DEFAULT NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;
	  `
)
