package db

import "fmt"

// Errors
var (
	ErrVaultNotFound = &Error{"vault not found", nil}
	ErrNoVaultName   = &Error{"no vault name", nil}
)

// Vault represents a username-pass pair
type Vault struct {
	Tx   *Tx
	Name []byte
	Credentials []byte
}

func (v *Vault) bucket() []byte {
	return []byte("vault")
}

func (v *Vault) get() ([]byte, error) {
	text := v.Tx.Bucket(v.bucket()).Get(v.Name)
	if text == nil {
		return nil, ErrVaultNotFound
	}

	return text, nil
}

// Load retrieves a vault from the database.
func (v *Vault) Load() error {
	name, err := v.get()
	if err != nil {
		return err
	}
	fmt.Printf("name: %s\n", name)
	credentials, err := v.get()
	fmt.Printf("credentials %s", credentials)
	if err != nil {
		return err
	}

	v.Credentials = credentials
	v.Name = name

	return nil
}

// Save commits the Vault to the database.
func (v *Vault) Save() error {
	if len(v.Name) == 0 {
		return ErrNoVaultName
	}

	return v.Tx.Bucket(v.bucket()).Put(v.Name, v.Credentials)
}
