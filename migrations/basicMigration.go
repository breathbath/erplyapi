package migrations

type BasicItem struct{}

func (bm BasicItem) GetPostQueryCallback() func() error {
	return func() error { return nil }
}
