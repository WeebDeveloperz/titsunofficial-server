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
  "net/http"
)

func Routes(route *gin.Engine) {
	nd := route.Group("/notes")
	{
		nd.GET("/list", func (ctx *gin.Context) {
	    //// TODO: add functionality to filter results
	    //clients, err := getClients(nil)
	    //if err != nil {
	    //	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	    //	log.Printf("ERROR: Failed to read documents: %v\n", err.Error())
	    //	return
	    //}

	    ctx.JSON(http.StatusOK, dirToJSON("data"))
		})
	}
}
