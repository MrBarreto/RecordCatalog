package drivers

import (
	"net/http"

	"github.com/MrBarreto/RecordCatalog/src/core/ports"
	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	access ports.RecordService
}

func NewRestHandler(RecordService ports.RecordService) *RestHandler {
	return &RestHandler{
		access: RecordService,
	}
}

func (r *RestHandler) GetAlbums(c *gin.Context) {
	albums, err := r.access.GetAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found any album"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func (r *RestHandler) GetAlbumsByArtist(c *gin.Context) {
	artist := c.Param("artist")
	albums, err := r.access.GetAlbumsByArtist(artist)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found any album of this artist"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func (r *RestHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := r.access.GetAlbumByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (r *RestHandler) GetAvailableArtists(c *gin.Context) {
	artists, err := r.access.GetAvailableArtists()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Artists not found"})
		return
	}
	c.JSON(http.StatusOK, artists)
}
