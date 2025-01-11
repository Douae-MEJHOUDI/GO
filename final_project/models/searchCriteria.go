package models

type SearchCriteria struct {
	Title    string   `json:"title,omitempty"`
	Author   string   `json:"author,omitempty"`
	Genres   []string `json:"genres,omitempty"`
	MinPrice *float64 `json:"min_price,omitempty"`
	MaxPrice *float64 `json:"max_price,omitempty"`
	InStock  *bool    `json:"in_stock,omitempty"`
}

func (sc *SearchCriteria) IsEmpty() bool {
	return sc.Title == "" &&
		sc.Author == "" &&
		len(sc.Genres) == 0 &&
		sc.MinPrice == nil &&
		sc.MaxPrice == nil &&
		sc.InStock == nil
}
