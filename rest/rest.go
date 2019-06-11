//Package rest defines the routes of the API
package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//RunAPIWithHandler runs the API and takes the host address as a parameter
func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
	}))
		
	r.GET("/artists", h.GetArtists)
	r.GET("/artists/:name", h.GetArtist)
	r.GET("/albums", h.GetAlbums)
	r.GET("/albums/:id", h.GetAlbum)
	r.GET("/tracks", h.GetTracks)
	return r.Run(address)
}

//RunAPI runs the API
func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}
