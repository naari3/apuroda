package handlers

import (
	"apuroda/stores"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CheckSession CheckSession
func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		strID := session.Get("user_id")
		if strID != nil {
			id, err := uuid.Parse(session.Get("user_id").(string))
			if err == nil {
				user, err := stores.UserStore.GetByID(id)
				if err == nil {
					c.Set("user", user)
				}
				c.Next()
				return
			}
		}
	}
}
