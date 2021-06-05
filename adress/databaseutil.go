package adress

import (
	"crypto/ecdsa"
	"database/sql"
	"math/big"
)

var wal []Wallet

func addwallet(priv, pub, puby string, db *sql.DB) (err error) {
	statement, err := db.Prepare(Add_key)
	ab, err := statement.Exec(priv, pub, puby)
	_ = ab

	return err

}

func listwallet(db *sql.DB) (wal []Wallet, err error) {
	wal = wal[:0]
	rows, err := db.Query(Ret_key)
	if err != nil {
		return
	}

	var a int
	var b, c, d string
	for rows.Next() {
		var wallet Wallet
		err = rows.Scan(&a, &b, &c, &d)
		dd := new(big.Int)
		wallet.privatekey = &ecdsa.PrivateKey{}
		wallet.privatekey.D, _ = dd.SetString(b, 10)
		wallet.privatekey.X, _ = dd.SetString(c, 10)
		wallet.privatekey.Y, _ = dd.SetString(d, 10)
		wal = append(wal, wallet)
	}
	return wal, err

}
