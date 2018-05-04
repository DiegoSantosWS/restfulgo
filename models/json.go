package models

/**
import(
	"encoding/json"
)


func NewClients(cli string) (*Clients, error) {
	cli = byte(cli)
	b := &Clients{}
	return b, json.Unmarshal([]byte(cli), b)
}

func NewAddress(address string) (*AddressClients, error) {
	a := &Address{}
	address = []byte(address)
	return a, json.Unmarshal([]byte(address), a)
}
**/

