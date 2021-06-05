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

	err := addwallet(wall.privatekey.D.String(), wall.privatekey.X.String(), wall.privatekey.Y.String(), db)
	if err != nil {
		println(err)
	}
	listwallet(db)
}

func (wall *Wallet) Listwallet(db *sql.DB) {
	wal, err := listwallet(db)
	if err != nil {
		fmt.Println("Erro")
	}
	for i, j := range wal {
		fmt.Printf("%d Public-Key : %x \n", i, j.privatekey.D)
	}
}
