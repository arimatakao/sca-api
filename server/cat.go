package server

import (
	"log"
	"net/http"

	"github.com/arimatakao/sca-api/external"
	"github.com/arimatakao/sca-api/server/storager"
	"github.com/gin-gonic/gin"
)

func (s *server) CatList(c *gin.Context) {
	cats, err := s.db.GetAllCats()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(cats) == 0 {
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusOK, cats)
}

func (s *server) SpecificCat(c *gin.Context) {
	catId := c.Param("id")
	if catId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad cat id",
		})
		return
	}

	cat, err := s.db.GetSpecificCat(catId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cat)
}

func (s *server) CreateCat(c *gin.Context) {
	var cat storager.Cat

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	isBreedAllow, err := external.IsBreedAllowed(cat.Breed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !isBreedAllow {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "breed is not allowed",
		})
		return
	}

	if err := s.db.CreateCat(cat); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (s *server) UpdateCat(c *gin.Context) {
	var catFields storager.Cat
	catId := c.Param("id")

	if catId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad cat id",
		})
		return
	}

	if err := c.ShouldBindJSON(&catFields); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	isBreedAllow, err := external.IsBreedAllowed(catFields.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !isBreedAllow {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "breed is not allowed",
		})
		return
	}

	if err := s.db.UpdateCat(catId, catFields); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (s *server) DeleteCat(c *gin.Context) {

	catId := c.Param("id")

	if catId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad cat id",
		})
		return
	}

	if err := s.db.DeleteCat(catId); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
