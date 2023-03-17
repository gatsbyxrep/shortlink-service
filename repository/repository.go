package repository

type IRepository interface {
	PushOriginalAndShort(original, short string) error
	// Returns original string or nil and error if string doesnt exist in rep
	GetByShortLink(short string) (string, error)
	Close()
}
