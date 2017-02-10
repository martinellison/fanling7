package persist

// A `PageMeta` describes a Page independently of the page type.
type PageMeta struct {
	Type  string
	Ident string
}

type Persistor interface {
	Open(path string)
	Close()
	GetPageList() (paths []string)
	ReadPage(ident string) (meta PageMeta, detail []byte, found bool)
	ReadPageFromPath(path string) (meta PageMeta, detail []byte, found bool)
	WritePage(meta PageMeta, detail []byte)
}
