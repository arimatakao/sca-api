package server

import (
	"database/sql"
	"net/http"

	"github.com/arimatakao/sca-api/server/storager"
	"github.com/gin-gonic/gin"
)

func (s *server) MissionsList(c *gin.Context) {
	missions, err := s.db.GetAllMissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(missions) == 0 {
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusOK, missions)
}

func (s *server) SpecificMission(c *gin.Context) {
	missionId := c.Param("id")
	if missionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad mission id",
		})
		return
	}

	mission, err := s.db.GetSpecificMission(missionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, mission)

}

func (s *server) CreateMission(c *gin.Context) {
	var mission storager.Mission

	if err := c.ShouldBindBodyWithJSON(&mission); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(mission.Targets) < 1 ||
		len(mission.Targets) > 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "too much mission targets",
		})
		return
	}

	_, err := s.db.GetSpecificCat(mission.AssigneeId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "assignee id is not exist",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	for _, target := range mission.Targets {
		createdTargetId, err := s.db.CreateMissionTarget(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		mission.AddCreatedTargetId(createdTargetId)
	}

	if err = s.db.CreateMission(mission); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (s *server) UpdateMission(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}

func (s *server) DeleteMission(c *gin.Context) {
	missionId := c.Param("id")
	if missionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad mission id",
		})
		return
	}

	mission, err := s.db.GetSpecificMission(missionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if mission.AssigneeId != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cat assigned on mission, it cannot be deleted",
		})
		return
	}

	if err = s.db.DeleteMission(missionId); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (s *server) MissionTargetList(c *gin.Context) {
	missionId := c.Param("id")
	if missionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad mission id",
		})
		return
	}

	targets, err := s.db.GetMissionTargets(missionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, targets)
}

func (s *server) CreateMissionTarget(c *gin.Context) {
	missionId := c.Param("id")

	if missionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad mission id",
		})
		return
	}

	var target storager.Target

	if err := c.ShouldBindBodyWithJSON(&target); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	targetId, err := s.db.CreateMissionTarget(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = s.db.AddTargetToMission(targetId, missionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (s *server) UpdateMissionTarget(c *gin.Context) {
	missionId := c.Param("id")
	targetId := c.Param("idTarget")
	if missionId == "" || targetId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad mission id or target id",
		})
		return
	}

	mission, err := s.db.GetSpecificMission(missionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "mission not exist",
		})
		return
	}

	if !mission.IsTargetExist(targetId) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Target is not exist on this mission",
		})
		return
	}

	if mission.IsComplete {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "changes in completed mission is not allowed",
		})
		return
	}

	target, err := s.db.GetTarget(targetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if target.IsComplete {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "changes in completed target is not allowed",
		})
		return
	}

	err = s.db.UpdateMissionTarget(targetId, storager.Target{
		Name:       target.Name,
		Country:    target.Country,
		Notes:      target.Notes,
		IsComplete: target.IsComplete,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
