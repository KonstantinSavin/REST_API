package sqldb

import (
	"effective-mobile/music-lib/internal/model"
	"fmt"
	"strings"
)

func (r SongRep) FilterCouplets(stp model.SongTextPagination) (*model.PaginatedText, error) {
	r.logger.Debugf("SongRep FilterSongs")

	text, err := r.GetSongTextByID(stp)
	if err != nil {
		return nil, err
	}

	couplets := formatCouplets(text)
	if couplets == nil {
		return nil, nil
	}

	var res model.PaginatedText
	res.Сouplets = couplets

	return &res, err
}

func (r SongRep) GetSongTextByID(stp model.SongTextPagination) (string, error) {
	r.logger.Debugf("SongRep GetSongTextByID")

	q := "SELECT text FROM songs WHERE id = $1"

	r.logger.Infof(fmt.Sprintf("SQL Query: %s, $1 = %d", q, stp.ID))

	var text string

	if err := r.storage.db.QueryRow(
		q,
		stp.ID,
	).Scan(&text); err != nil {
		r.logger.Errorf("Ошибка SQL: %s", err)
		return "", err
	}

	return text, nil
}

func formatCouplets(text string) []string {
	couplets := strings.Split(strings.TrimSpace(text), "\n\n")

	formattedCouplets := make([]string, len(couplets))

	for i, v := range couplets {
		formattedCouplet := strings.ReplaceAll(v, "\n", "\n    ")
		title := fmt.Sprintf("Куплет %d:\n    ", i+1)
		formattedCouplets[i] = title + formattedCouplet
	}
	return formattedCouplets
}
