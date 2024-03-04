package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zenkimoto/vitals-server-api/internal/env"
	"github.com/zenkimoto/vitals-server-api/internal/util"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		if header == "" {
			log.Printf("Authorization header is missing.")
			c.String(401, "Unauthorized")
			c.Abort()
			return
		}

		if ar := strings.Split(header, "Bearer "); len(ar) == 2 {
			user, id, err := util.Parse(env.GetJWTKey(), ar[1])
			if err != nil {
				log.Print(err)
				c.String(401, "Unauthorized")
				c.Abort()
				return
			} else {
				c.Set("user", user)
				c.Set("id", id)
			}
		} else {
			log.Print("Can not parse Authorization header.")
			c.String(401, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}
