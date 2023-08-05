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

package auth

import (
	"net/http"
	"strings"
	"log"
	"errors"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	a := route.Group("/api/auth/login")
	{
		a.POST("/", func(ctx *gin.Context) {
			username := ctx.PostForm("username")
			password := ctx.PostForm("password")

			if len(strings.TrimSpace(username)) == 0 || len(strings.TrimSpace(password)) == 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is empty"})
				return
			}

			var u User
      err := db.Where("username = ?", username).First(&u).Error
			if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
				  ctx.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})
				} else {
				  ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
					log.Printf("Unexpected error while logging in: %v\n", err)
				}

				return
			}

			if u.Password == password {
        tk, err := newJWT(u.Username, u.Role)
				if err != nil {
					log.Printf("Error while generating JWT for %s: %v", u.Username, err.Error())
				  ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				} else {
			  	log.Printf("User %s logged in\n", u.Username)
			  	ctx.JSON(http.StatusOK, gin.H{"message": "logged in", "token": tk})
				}
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "password is incorrect"})
			}
		})
	}
}
