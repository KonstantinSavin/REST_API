package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func (s *Service) fetchInfoFromAPI(group, song string) (SongDetail, error) {
	s.logger.Debugf("Запрос на информацию по песне %s группы %s в API", song, group)

	params := url.Values{}
	params.Add("group", group)
	params.Add("song", song)
	urlAPI := fmt.Sprintf("%s?%s", s.apiUrl, params.Encode())

	s.logger.Debugf("Запрос %s", urlAPI)

	resp, err := http.Get(urlAPI)
	if err != nil {
		s.logger.Errorf("Произошла ошибка при запросе: %v\n", err)
		return SongDetail{}, err
	}
	defer resp.Body.Close()

	var songDetail SongDetail
	switch resp.StatusCode {
	case 200:
		body, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &songDetail); err != nil {
			fmt.Println("Ошибка парсинга ответа:", err)
			return SongDetail{}, err
		}

		s.logger.Infof("Информация о песне:\nДата выхода: %s\nТекс песни: %s\nСсылка: %s",
			songDetail.ReleaseDate,
			songDetail.Text,
			songDetail.Link)
	case 400:
		s.logger.Errorln("Ошибка: неверный запрос. Проверьте параметры `group` и `song`.")
	case 500:
		s.logger.Errorln("Ошибка сервера: попробуйте ещё раз позже.")
	default:
		s.logger.Errorf("Ошибка: %d\n", resp.StatusCode)
	}
	return songDetail, nil
}
