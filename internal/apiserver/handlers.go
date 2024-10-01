package apiserver

import (
	"effective-mobile/music-lib/internal/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *server) handlerAddSong(c *gin.Context) {
	song := model.Song{}
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := srv.storage.Song().CreateSong(&song)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": "song was added", "song": song})
}

func (srv *server) handlerDeleteSong(c *gin.Context) {
	songID := c.Param("id")

	err := srv.storage.Song().DeleteSong(songID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("song with id = %s deleted", songID)})
}

func (srv *server) handlerUpdateSong(c *gin.Context) {
	songID := c.Param("id")

	var newSong model.Song
	if err := c.ShouldBindJSON(&newSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err := srv.storage.Song().UpdateSong(songID, &newSong)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("song with id = %s has been updated", songID), "song": song})
}
