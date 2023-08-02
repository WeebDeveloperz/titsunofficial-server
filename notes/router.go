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

package notes

import (
  "github.com/gin-gonic/gin"
	"log"
	"encoding/json"
  "net/http"
)

func Routes(route *gin.Engine) {
	s := route.Group("/subjects")
	{
		s.GET("/", func(ctx *gin.Context) {
			var subjects []Subject

			// TODO: handle error
			res := db.Find(&subjects)
			log.Printf("Read all subjects from DB: %v", res)

			ctx.JSON(http.StatusOK, gin.H{"data": subjects})
		})

		s.POST("/", func(ctx *gin.Context) {
			var s Subject
      json.Unmarshal([]byte(ctx.PostForm("data")), &s)
			log.Println(ctx.PostForm("data"))

			// TODO: handle error
			res := db.Create(&s)
			log.Printf("Saved new subject to DB: %v", res)

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		s.PUT("/", func(ctx *gin.Context) {
			var s Subject
      json.Unmarshal([]byte(ctx.PostForm("data")), &s)

			// TODO: handle error
			res := db.Save(&s)
			log.Printf("Updated subject in DB: %v", res)

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		s.DELETE("/", func(ctx *gin.Context) {
			var s Subject
      json.Unmarshal([]byte(ctx.PostForm("data")), &s)

			// TODO: handle error
			res := db.Delete(&s)
			log.Printf("Deleted subject from DB: %v", res)

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	f := route.Group("/files")
	{
		f.POST("/", func(ctx *gin.Context) {
			var f File
      json.Unmarshal([]byte(ctx.PostForm("data")), &f)

			// TODO: handle error
			res := db.Create(&f)
			log.Printf("Saved new file to DB: %v", res)

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		f.GET("/all", func(ctx *gin.Context) {
			var files []File

			// TODO: handle error
			res := db.Find(&files)
			log.Printf("Read all files from DB: %v", res)

			ctx.JSON(http.StatusOK, gin.H{"data": files})
		})
	}
}
