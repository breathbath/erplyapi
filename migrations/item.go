package migrations

type Item interface {
	GetId() string
	GetQueries() []Query
	GetPostQueryCallback() func() error
}
