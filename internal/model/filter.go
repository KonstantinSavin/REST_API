package model

type Filter struct {
	Page    *int    `json:"page"`
	PerPage *int    `json:"per_page"`
	ID      *string `json:"id"`
	Name    *string `json:"song"`
	Group   *string `json:"group"`
}

type FilteredSongs struct {
	Songs []*Song `json:"songs"`
}

func (f Filter) Update() Filter {
	if f.Page == nil {
		page := 1
		f.Page = &page
	}
	if f.PerPage == nil {
		perPage := 10
		f.PerPage = &perPage
	}

	return Filter{
		Page:    f.Page,
		PerPage: f.PerPage,
		ID:      f.ID,
		Name:    f.Name,
		Group:   f.Group,
	}
}
