package setup

var (
	CreateDatabase  = `CREATE DATABASE IF NOT EXISTS alpha_transaction DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;`
	UseAlphaTransaction    = `USE alpha_transaction`
	SetTimeZone     = "SET time_zone = '+00:00';"
	DropDB          = `DROP DATABASE IF EXISTS alpha_transaction;`
	SQLMode           = "SET SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO'"
	ClearTransactionTable    = "DELETE FROM transactions;"

	CreateTransactionTable = `CREATE TABLE IF NOT EXISTS transactions (
		id varchar(255) NOT NULL UNIQUE,
		created_at timestamp NULL DEFAULT NULL,
		updated_at timestamp NULL DEFAULT NULL,
		deleted_at timestamp NULL DEFAULT NULL,
		memo varchar(255)  NULL,
		wallet_id varchar(100)  NULL,
		payment_type varchar(100)  NULL,
		customer varchar(100)  NULL,
		customer_id varchar(100)  NULL,
		tx_ref varchar(100)  NULL,
		amount  varchar(255) NOT NULL,
		flw_ref varchar(100)  NULL,
		currency varchar(100)  NULL,
		status varchar(100)  NULL,
		message  varchar(255)  NULL,
		card_last_four_digit varchar(100)  NULL,
		card_type varchar(100)  NULL,
		third_party_id varchar(100)  NULL,
		source varchar(255) NULL,
		trans_type tinyint(1)  NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;
	  `
)