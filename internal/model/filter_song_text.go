package model

type SongTextPagination struct {
	Page          *int           `json:"page"`
	PerPage       *int           `json:"per_page"`
	ID            *int           `json:"id"`
	Text          *string        `json:"text"`
	PaginatedText *PaginatedText `json:"paginated_text"`
}

type PaginatedText struct {
	Сouplets []string `json:"couplets"`
}

func (stp SongTextPagination) Update() SongTextPagination {
	if stp.Page == nil {
		page := 1
		stp.Page = &page
	}
	if stp.PerPage == nil {
		perPage := 3
		stp.PerPage = &perPage
	}

	return SongTextPagination{
		Page:    stp.Page,
		PerPage: stp.PerPage,
		ID:      stp.ID,
	}
}
