package search

type defaultMacher struct{}

func init() {
	var matcher defaultMacher
	Register("default", matcher)
}
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
