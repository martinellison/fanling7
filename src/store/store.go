package store

type Store interface {
	Add(storable Storable)
	Changed(storable Storable)
	Get(ident string) (storable Storable, found bool)
	Close()
}

type Storable interface {
	Ident() (ident string)
}

type StoreHelper interface{}
