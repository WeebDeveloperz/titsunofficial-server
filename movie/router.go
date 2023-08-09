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

package movie

import (
	"log"
	"net/http"

	"github.com/WeebDeveloperz/titsunofficial-server/auth"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	m := route.Group("/api/movie")
	{
		m.GET("/", auth.Authorize("read"), func(ctx *gin.Context) {
			var movies []Movie

		  res := db.Find(&movies)
			if res.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			  log.Printf("Unexpected error while reading all movies from DB: %v\n", res.Error)
				return
			}

			log.Printf("Read all movies from DB.\n")
			ctx.JSON(http.StatusOK, gin.H{"data": movies})
		})

		m.POST("/", func(ctx *gin.Context) {
			var m Movie
			ctx.Bind(&m)

			res := db.Create(&m)
			if res.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			  log.Printf("Unexpected error while saving new movie to DB: %v\n", res.Error)
				return
			}

			log.Printf("Saved new movie to DB: %d\n", m.ID)
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
}
