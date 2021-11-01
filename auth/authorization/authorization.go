package authorization

import (
	"budget-calendar/database"
	"github.com/gin-gonic/gin"
)

func ValidSession(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := db.SessionStore.Get(c.Request, "session")
		if err != nil {
			c.AbortWithStatusJSON(500, "The server was unable to retrieve this session")
			return
		}

		if session.ID == "" {
			c.AbortWithStatusJSON(401, "This user has no current session. Use of this endpoint is thus unauthorized")
			return
		}
	}
}