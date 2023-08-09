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
	"gorm.io/gorm"
	"github.com/WeebDeveloperz/titsunofficial-server/database"
)

var db *gorm.DB
func Init() {
	db = database.DB
	db.AutoMigrate(&Movie{})
}

type Movie struct {
	ID        uint   `json:"ID"`
	Title     string `json:"movie_title"`
	Name      string `json:"submitter_name"`
	CollegeID string `json:"submitter_college_id"`
}
