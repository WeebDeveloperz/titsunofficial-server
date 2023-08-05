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
	"github.com/google/uuid"
	"log"
	"encoding/json"
  "net/http"
	"os"
)

func Routes(route *gin.Engine) {
	s := route.Group("/api/subjects")
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
			var files []File
			db.Where("subject_id = ?", s.ID).Find(&files)

			db.Delete(&files)
			log.Printf("Deleted all files for subject \"%s\" from DB: %v", s.SubjectCode, files)

			// TODO: handle error
			res := db.Delete(&s)
			log.Printf("Deleted subject from DB: %v", res)

			for _, f := range files {
			  fp := dataDir + f.FilePath
			  err := os.Remove(fp)
			  if err != nil {
			  	// TODO: check what error it is
			  	log.Printf("Error while deleting files: %v\n", err.Error())
			    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			  }
			}
			log.Printf("Deleted all files for subject \"%s\" from filesystem: %v", s.SubjectCode, files)

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	f := route.Group("/api/files")
	{
		f.GET("/", func(ctx *gin.Context) {
			var files []File

			// TODO: handle error
			res := db.Find(&files)
			log.Printf("Read all files from DB: %v\n", res)

			ctx.JSON(http.StatusOK, gin.H{"data": files})
		})

		f.POST("/", func(ctx *gin.Context) {
			var f File
      json.Unmarshal([]byte(ctx.PostForm("data")), &f)

			file, err := ctx.FormFile("file")
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while getting FormFile: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

      f.FilePath = uuid.New().String() + ".pdf"
      err = ctx.SaveUploadedFile(file, dataDir + f.FilePath)
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while saving uploaded file: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

			// TODO: handle error
			res := db.Create(&f)
			log.Printf("Saved new file to DB: %v", res)

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		f.DELETE("/", func(ctx *gin.Context) {
			var f File
      json.Unmarshal([]byte(ctx.PostForm("data")), &f)

			fp := dataDir + f.FilePath

			// TODO: handle error
			res := db.Delete(&f)
			log.Printf("Deleted file from DB: %v\n", res)

			err := os.Remove(fp)
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while deleting file: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
}
