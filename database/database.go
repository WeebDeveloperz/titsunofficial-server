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

package database

import (
	"gorm.io/gorm"
  "gorm.io/driver/mysql"
	"os"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("dsn")), &gorm.Config{})

	if err != nil {
		log.Printf("Error while connecting to database: %v\n", err.Error())
		os.Exit(1)
	}
}
