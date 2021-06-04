package adress

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"database/sql"
	"fmt"
)

type Wallet struct {
	privatekey *ecdsa.PrivateKey
}

func (wall Wallet) PrintPriv() {
	fmt.Printf("\n %x", wall.privatekey.D)

}

func (wall Wallet) Read(b []byte) (int, error) {
	rand.Read(b)
	return len(b), nil

}
func (wall *Wallet) Genwallet(db *sql.DB) {
	wall.privatekey, _ = ecdsa.GenerateKey(elliptic.P224(), wall)
	fmt.Println(wall.privatekey.X)
	statement, err := db.Prepare("INSERT INTO WALLET (privatekey,publickey) VALUES (?,?)")
	statement.Exec(fmt.Sprintf("%x", wall.privatekey.D), fmt.Sprintf("%x", wall.privatekey.PublicKey.X))

	if err != nil {
		fmt.Println("there is an error")
	}

}
