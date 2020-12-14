package db
// TODO auth will be added later, until then this is useless
// Errors
var (
	ErrUserNotFound = &Error{"vault not found", nil}
	ErrNoUserName   = &Error{"no vault name", nil}
)

// User represents a username-pass pair
type User struct {
	Tx   *Tx
	Name     []byte
	Username []byte
	Password []byte
}

func (v *User) bucket() []byte {
	return []byte("vault")
}

func (v *User) get() ([]byte, error) {
	text := v.Tx.Bucket(v.bucket()).Get(v.Name)
	if text == nil {
		return nil, ErrUserNotFound
	}

	return text, nil
}

// Load retrieves a vault from the database.
func (v *User) Load() error {
	username, err := v.get()
	if err != nil {
		return err
	}
	password, err := v.get()
	if err != nil {
		return err
	}

	v.Username = username
	v.Password = password

	return nil
}

// Save commits the User to the database.
func (v *User) Save() error {
	if len(v.Name) == 0 {
		return ErrNoUserName
	}

	return v.Tx.Bucket(v.bucket()).Put(v.Username, v.Password)
}

