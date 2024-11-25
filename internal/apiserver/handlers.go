package apiserver

import (
	_ "effective-mobile/music-lib/docs"
	"effective-mobile/music-lib/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddSong godoc
//
// @Summary      AddSong
// @Description  add new song to library
// @Tags         songs
// @ID           add-song
// @Accept       appcication/json
// @Produce      appcication/json
// @Param        input  body            model.Song  true  "add song"
// @Success      200          {object}  model.Song
// @Router       /add [post]
func (srv *server) handlerAddSong(c *gin.Context) {
	srv.logger.Debug("Handler AddSong")

	song := model.Song{}
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enrichedSong, err := srv.service.AddSong(song)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": "song was added", "song": enrichedSong})
}

// DeleteSong godoc
//
// @Summary      DeleteSong
// @Description  delete song from library
// @Tags         songs
// @ID           delete-song
// @Accept       appcication/json
// @Produce      appcication/json
// @Param        id   path   string     true  "song id"
// @Success      204  {object}  model.Song
// @Router       /delete/{id} [delete]
func (srv *server) handlerDeleteSong(c *gin.Context) {
	srv.logger.Debug("Handler DeleteSong")

	songID := c.Param("id")

	err := srv.service.DeleteSong(songID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("song with id = %s deleted", songID)})
}

// UpdateSong godoc
//
// @Summary      UpdateSong
// @Description  update song from library
// @Tags         songs
// @ID           update-song
// @Accept       appcication/json
// @Produce      appcication/json
// @Param        id           path                  string    true  "song id"
// @Param        input  body            model.Song  true    "delete song"
// @Success      200          {object}  model.Song
// @Router       /update/{id} [patch]
func (srv *server) handlerUpdateSong(c *gin.Context) {
	srv.logger.Debug("Handler UpdateSong")

	songID := c.Param("id")

	var newSong model.EnrichedSong
	if err := c.ShouldBindJSON(&newSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err := srv.service.UpdateSong(songID, &newSong)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("song with id = %s has been updated", songID), "song": song})
}

// GetSongs godoc
//
// @Summary      GetSongs
// @Description  get songs with filtration and pagination
// @Tags         songs
// @ID           get-songs
// @Accept       appcication/json
// @Produce      appcication/json
// @Param        input  body  storage.Filter  true  "filter"
// @Success      200          {array}         storage.Filter
// @Router       /songs [post]
func (srv *server) handlerGetSongs(c *gin.Context) {
	srv.logger.Debug("Handler GetSongs")

	input := model.Filter{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := input.Update()
	songs, hasNextPagge, err := srv.service.GetSongs(&filter)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.Writer.Header().Set("Pagination-Page", strconv.Itoa(*filter.Page))
	c.Writer.Header().Set("Pagination-Limit", strconv.Itoa(*filter.PerPage))

	c.Writer.Header().Set("Has-Next-Page", strconv.FormatBool(hasNextPagge))

	c.JSON(
		http.StatusOK,
		model.FilteredSongs{Songs: songs},
	)
}
