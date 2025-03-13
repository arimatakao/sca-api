package server

import (
	"github.com/arimatakao/sca-api/config"
	"github.com/arimatakao/sca-api/server/storager"
	"github.com/gin-gonic/gin"
)

type server struct {
	r  *gin.Engine
	db storager.DbStorager
}

func New() server {
	gin.SetMode(gin.DebugMode)
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
	base.DELETE(catsPath+"/:id", s.DeleteCat)

	missionsPath := "/missions"
	base.GET(missionsPath, s.MissionsList)
	base.GET(missionsPath+"/:id", s.SpecificMission)
	base.POST(missionsPath, s.CreateMission)
	base.PUT(missionsPath+"/:id", s.UpdateMission)
	base.DELETE(missionsPath+"/:id", s.DeleteMission)

	base.GET(missionsPath+"/:id/targets", s.MissionTargetList)
	base.POST(missionsPath+"/:id/targets", s.CreateMissionTarget)
	base.PUT(missionsPath+"/:id/targets/:idTarget", s.UpdateMissionTarget)

	db, err := storager.NewPostgreSQL(config.App.DbUrl)
	if err != nil {
		return err
	}
	s.db = db
	return s.r.Run(":" + config.App.Port)
}
