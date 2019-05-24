//Package rest defines the routes of the API
//Here are defined the handlers of each route
package rest

import (
	"musicapi/dblayer"
	"net/http"

	"github.com/gin-gonic/gin"
)

//HandlerInterface defines the interface of the handler
type HandlerInterface interface {
	GetArtists(c *gin.Context)
	GetArtist(c *gin.Context)
	GetAlbums(c *gin.Context)
	GetAlbum(c *gin.Context)
	GetTracks(c *gin.Context)
}

//Handler hold a reference of the database connection
type Handler struct {
	db dblayer.DBLayer
}

//NewHandler creates a new handler object
func NewHandler() (HandlerInterface, error) {
	return NewHandlerWithParams("sqlite3", "music.db")
}

//NewHandlerWithParams creates a new handler object based upon the paramaters
func NewHandlerWithParams(dbtype, conn string) (HandlerInterface, error) {
	db, err := dblayer.NewDB(dbtype, conn)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}

//GetArtists handler for the route /artists
func (h *Handler) GetArtists(c *gin.Context) {
	if h.db == nil {
		return
	}
	artists, err := h.db.GetAllArtists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artists)
}

//GetArtist handler for the route /artists/:name
func (h *Handler) GetArtist(c *gin.Context) {
	if h.db == nil {
		return
	}
	artists, err := h.db.GetArtist(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artists)
}

//GetAlbums handler for the route /albums
func (h *Handler) GetAlbums(c *gin.Context) {
	if h.db == nil {
		return
	}
	albums, err := h.db.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

//GetAlbum handler for the route /albums/:id
func (h *Handler) GetAlbum(c *gin.Context) {
	if h.db == nil {
		return
	}
	albums, err := h.db.GetAlbum(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

//GetTracks handler for the route /tracks
func (h *Handler) GetTracks(c *gin.Context) {
	if h.db == nil {
		return
	}
	tracks, err := h.db.GetAllTracks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tracks)
}
