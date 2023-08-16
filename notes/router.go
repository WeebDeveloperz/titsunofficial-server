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
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/WeebDeveloperz/titsunofficial-server/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine) {
	s := route.Group("/api/subjects")
	{
		s.GET("/", func(ctx *gin.Context) {
			var subjects []Subject

			branch := ctx.Query("branch")
			semester := ctx.Query("semester")

			// I wanna kill myself
			var res *gorm.DB
			if branch == "" && semester == "" {
			  res = db.Find(&subjects)
			} else if branch == "" && semester != "" {
			  res = db.Where("semester = ?", semester).Find(&subjects)
			} else if branch != "" && semester == "" {
			  res = db.Where("branch = ?", branch).Find(&subjects)
			} else {
			  res = db.Where("branch = ? and semester = ?", branch, semester).Find(&subjects)
			}

			if res.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			  log.Printf("Unexpected error while reading subects from DB: %v\n", res.Error)
				return
			}

			log.Printf("Read all subjects from DB.\n")
			ctx.JSON(http.StatusOK, gin.H{"data": subjects})
		})

		s.POST("/", auth.Authorize("write"), func(ctx *gin.Context) {
			var s Subject
      json.Unmarshal([]byte(ctx.PostForm("data")), &s)

			file, err := ctx.FormFile("file")
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while getting FormFile: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

      s.ImagePath = uuid.New().String() + ".webp"
      err = ctx.SaveUploadedFile(file, subImgDir + s.ImagePath)
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while saving uploaded file: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

			username, _ := ctx.Get("username")
			s.CreatedBy = username.(string)

			res := db.Create(&s)
			if res.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			  log.Printf("Unexpected error while saving new subect to DB: %v\n", res.Error)
				return
			}

			log.Printf("Saved new subject to DB: %d\n", s.ID)
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		s.PUT("/", auth.Authorize("write"), func(ctx *gin.Context) {
			var s Subject
      json.Unmarshal([]byte(ctx.PostForm("data")), &s)

			username, _ := ctx.Get("username")
			s.UpdatedBy = username.(string)

			res := db.Save(&s)
			if res.Error != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			  log.Printf("Unexpected error while updating subect in DB: %v\n", res.Error)
				return
			}

			log.Printf("Updated subject %d in DB\n", s.ID)
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		s.DELETE("/", auth.Authorize("delete"), func(ctx *gin.Context) {
			var s Subject
      json.Unmarshal([]byte(ctx.PostForm("data")), &s)

			var files []File
			db.Where("subject_id = ?", s.ID).Find(&files)

			db.Delete(&files)
			log.Printf("Deleted all files for subject \"%s\" from DB.\n", s.SubjectCode)

			res := db.Delete(&s)
			if res.Error != nil {
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			  log.Printf("Failed to delete from DB: %d\n", s.ID)
				return
			}

			username, _ := ctx.Get("username")
			log.Printf("User \"%s\" Deleted subject from DB: %d\n", username, s.ID)

			for _, f := range files {
			  fp := notesDir + f.FilePath
			  err := os.Remove(fp)
			  if err != nil {
			  	// TODO: check what error it is
			  	log.Printf("Error while deleting files: %v\n", err.Error())
			    ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			  }
			}

			log.Printf("Deleted all files for subject \"%s\" from filesystem.\n", s.SubjectCode)
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}

	f := route.Group("/api/files")
	{
		f.GET("/", func(ctx *gin.Context) {
			var files []File
			subId := ctx.Query("sub_id")

			// I wanna kill myself
			var res *gorm.DB
			if subId == ""  {
			  res = db.Preload("Subject").Find(&files)
			} else {
			  res = db.Preload("Subject", "id = ?", subId).Find(&files)
			}

			if res.Error != nil {
  			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			  log.Printf("Error while reading all files from DB: %v\n", res.Error)
				return
			}

			log.Printf("Read all files from DB.\n")

			filesFiltered := []File{}
			for _, i := range files {
				if i.Subject.ID != 0 {
					filesFiltered = append(filesFiltered, i)
				}
			}

			ctx.JSON(http.StatusOK, gin.H{"data": filesFiltered})
		})

		f.POST("/", auth.Authorize("write"), func(ctx *gin.Context) {
			var f File
      json.Unmarshal([]byte(ctx.PostForm("data")), &f)

			file, err := ctx.FormFile("file")
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while getting FormFile: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

      f.FilePath = uuid.New().String() + ".pdf"
      err = ctx.SaveUploadedFile(file, notesDir + f.FilePath)
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while saving uploaded file: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

			username, _ := ctx.Get("username")
			f.CreatedBy = username.(string)

			res := db.Create(&f)
			if res.Error != nil {
			  ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				log.Printf("Error while adding new file to DB: %v\n", res.Error)
				return
			}
			log.Printf("Saved new file to DB.")

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		f.DELETE("/", auth.Authorize("delete"), func(ctx *gin.Context) {
			var f File
      json.Unmarshal([]byte(ctx.PostForm("data")), &f)

			fp := notesDir + f.FilePath

			res := db.Delete(&f)
			if res.Error != nil {
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			  log.Printf("Failed to delete file from DB: %v\n", res.Error)
				return
			}

			username, _ := ctx.Get("username")
			log.Printf("User \"%s\" deleted file %d from DB\n", username, f.ID)

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
