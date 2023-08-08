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

package other

import (
	"log"
	"net/http"

	"github.com/WeebDeveloperz/titsunofficial-server/auth"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	o := route.Group("/api/other")
	{
		o.POST("/set-index-post", auth.Authorize("write"), func(ctx *gin.Context) {
			file, err := ctx.FormFile("file")
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while getting index.webp FormFile: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

			log.Println(dataDir + "index.webp")

      err = ctx.SaveUploadedFile(file, dataDir + "index.webp")
			if err != nil {
				// TODO: check what error it is
				log.Printf("Error while updating index.webp: %v\n", err.Error())
			  ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			}

			username, _ := ctx.Get("username")
			log.Printf("User\"%s\" updated index.webp.\n", username.(string))

			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
}
