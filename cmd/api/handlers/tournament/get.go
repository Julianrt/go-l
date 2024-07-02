package tournament

import (
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-l/cmd/api/core"
)

// GetTournament godoc
// @Summary Get a tournament
// @Description get a tournament with id param
// @Tags tournaments
// @Accept  json
// @Produce  json
// @Param id path string true "tournament id"
// @Success 200 {object} map[string]interface{} "tournament: domain.Tournament"
// @Failure 400 {object} map[string]interface{} "error: string"
// @Router /tournaments/:id [get]
func (h Handler) GetTournament(c *gin.Context) {
	tournamentIdParam := c.Param("id")
	tournament, err := h.TournamentService.Get(c, tournamentIdParam)
	if err != nil {
		core.RespondError(c, err)
		return
	}

	c.JSON(200, tournament)
}
