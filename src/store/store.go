package store

type Store interface {
	Add(storable Storable)
	Changed(storable Storable)
	Get(ident string) (storable Storable, found bool)
	GetByIndex(indexNum int) []Storable
	Flag(flagNum int, storable Storable)
	GetByFlagAndClear(flagNum int) map[string]Storable
	Close()
}

type Storable interface {
	Ident() (ident string)
	IndexKey(indexNum int) string
}

const MaxFlag = 1
const MaxIndex = 1
