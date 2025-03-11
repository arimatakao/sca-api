package server

import (
	"github.com/arimatakao/sca-api/config"
	"github.com/arimatakao/sca-api/server/storager"
	"github.com/gin-gonic/gin"
)

type server struct {
	r  *gin.Engine
	db storager.CatStorager
}

func New() server {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return server{
		r: router,
	}
}

func (s server) Run() error {
	base := s.r.Group(config.App.BasePath)

	catsPath := "/cats"
	base.GET(catsPath, s.CatList)
	base.GET(catsPath+"/:id", s.SpecificCat)
	base.POST(catsPath, s.CreateCat)
	base.PUT(catsPath+"/:id", s.UpdateCat)
	base.DELETE(catsPath, s.DeleteCat)

	// missionsPath := "/missions"
	// base.GET(missionsPath, MissionsList)
	// base.GET(missionsPath+"/:id", SpecificMission)
	// base.POST(missionsPath, CreateMission)
	// base.PUT(missionsPath+"/:id", UpdateMission)
	// base.DELETE(missionsPath, DeleteMission)

	// base.GET(missionsPath+"/:id/targets", MissionTargetList)
	// base.POST(missionsPath+"/:id/targets/:idTarget", CreateMissionTarget)
	// base.PUT(missionsPath+"/:id/targets/:idTarget", UpdateMissionTarget)

	db, err := storager.NewPostgreSQL(config.App.DbUrl)
	if err != nil {
		return err
	}
	s.db = db
	return s.r.Run(":" + config.App.Port)
}
