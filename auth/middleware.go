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

	"github.com/gin-gonic/gin"
)

func getPriorityByRoleString(r string) int {
	switch(r) {
	case "write":
		return 1
	case "delete":
		return 2
	default:
		return 0
	}
}

func Authorize(role string) gin.HandlerFunc {
  return func(ctx *gin.Context) {
	  tk := ctx.PostForm("token")

		claims, err := parseJWT(tk)
		if err != nil {
      ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
      return
		}

		ctx.Set("username", claims.Username)
		ctx.Set("role", claims.Role)

		if claims.Role == "admin" {
			ctx.Next()
			return
		}

	  if getPriorityByRoleString(role) > getPriorityByRoleString(claims.Role) {
      ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized to do this task"})
			return
	  }
  }
}
