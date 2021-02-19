package search

type defaultMacher struct{}

func init() {
	var matcher defaultMacher
	Register("default", matcher)
}
func (m defaultMacher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
