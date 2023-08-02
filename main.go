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
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
	d.ConnectToDB()
	n.Init() // shitty

  r := gin.New()

	n.Routes(r)

  r.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
  })

  r.Run(":6969")
}
