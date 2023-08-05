/* titsunofficial-server - Server for unofficial TIT&S website (github.com/WeebDeveloperz/titsunofficial)
 * Copyright (C) 2023  titsunofficial-server contributors

 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	_ "github.com/joho/godotenv/autoload"
  d "github.com/WeebDeveloperz/titsunofficial-server/database"
  n "github.com/WeebDeveloperz/titsunofficial-server/notes"
  a "github.com/WeebDeveloperz/titsunofficial-server/auth"
  "github.com/gin-gonic/gin"
  "net/http"
	"os"
)

func main() {
	d.ConnectToDB()
	n.Init() // shitty
	a.Init() // shitty

  r := gin.New()

	// cors
	r.Use(func (c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

    if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(204)
      return
    }

    c.Next()
	})

	n.Routes(r)
	a.Routes(r)

  r.GET("/api/ping", func(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
  })

  r.Run(":" + os.Getenv("PORT"))
}
