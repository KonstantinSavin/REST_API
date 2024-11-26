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
// @Summary      Add a new song to the library
// @Description  This endpoint allows to add a new song to the library.
// @Tags         songs
// @ID           add-song
// @Accept       application/json
// @Produce      application/json
// @Param        input  body            model.Song  true  "Song object to be added"
// @Success      201          {object}  model.EnrichedSong  "Successfully created song"
// @Failure      400          {object}  string  "Invalid input"
// @Failure      422          {object}  string  "Unprocessable entity"
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
// @Summary      Delete a song from the library
// @Description  This endpoint allows the user to delete a song by its ID from the library.
// @Tags         songs
// @ID           delete-song
// @Accept       application/json
// @Produce      application/json
// @Param        id   path   string     true  "ID of the song to be deleted"
// @Success      204   "Successfully deleted song"
// @Failure      422   {object}  string  "Unprocessable entity"
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
// @Summary      Update a song in the library
// @Description  This endpoint allows the user to update an existing song by its ID in the library.
// @Tags         songs
// @ID           update-song
// @Accept       application/json
// @Produce      application/json
// @Param        id    path   string              true  "ID of the song to be updated"
// @Param        input body    model.EnrichedSong  true  "New song data"
// @Success      200   {object}  model.EnrichedSong
// @Failure      400   {object}  string  "Invalid input"
// @Failure      422   {object}  string  "Unprocessable entity"
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
// @Summary      Retrieves a list of songs with filtering and pagination
// @Description  This endpoint allows the user to retrieve a filtered list of songs with pagination information.
// @Tags         songs
// @ID           get-songs
// @Accept       application/json
// @Produce      application/json
// @Param        input body model.Filter true "Filtering parameters for songs"
// @Success      200   {array}   model.EnrichedSong "A list of enriched songs"
// @Failure      400   {object}  string      "Bad request"
// @Failure      404   {object}  string      "Songs not found"
// @Failure      422   {object}  string      "Unprocessable entity"
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

	if len(songs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Songs not found"})
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

// GetCouplets godoc
//
// @Summary      Get the song text with pagination by couplets
// @Description  This endpoint retrieves the text of a song, broken down into couplets, with support for pagination.
// @Tags         songs
// @ID           get-couplets
// @Accept       application/json
// @Produce      application/json
// @Param        input  body  model.SongTextPagination  true  "Filters and pagination parameters"
// @Success      200          {object}         model.PaginatedText "Returns paginated couplets of the song"
// @Failure      400          {object}         string      "Bad request"
// @Failure      404          {object}         string      "Song not found"
// @Failure      422          {object}         string      "Unprocessable entity"
// @Router       /songtext/{id} [post]
func (srv *server) handlerGetCouplets(c *gin.Context) {
	srv.logger.Debug("Handler GetCouplets")

	input := model.SongTextPagination{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stp := input.Update()

	paginatedText, hasNextPage, err := srv.service.GetCouplets(&stp)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if len(paginatedText.Сouplets) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	c.Writer.Header().Set("Pagination-Page", strconv.Itoa(*stp.Page))
	c.Writer.Header().Set("Pagination-Limit", strconv.Itoa(*stp.PerPage))
	c.Writer.Header().Set("Has-Next-Page", strconv.FormatBool(hasNextPage))

	c.JSON(
		http.StatusOK,
		model.PaginatedText{Сouplets: paginatedText.Сouplets},
	)
}
