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
  "os"
	"time"
	"io/ioutil"
	"path/filepath"
)

type File struct {
  ModifiedTime time.Time `json:"ModifiedTime"`
  IsLink       bool      `json:"IsLink"`
  IsDir        bool      `json:"IsDir"`
  LinksTo      string    `json:"LinksTo"`
  Size         int64     `json:"Size"`
  Name         string    `json:"Name"`
  Path         string    `json:"Path"`
  Children     []*File   `json:"Children"`
}

func dirToJSON(path string) File {
  rootOSFile, _ := os.Stat(path)
  rootFile := toFile(rootOSFile, path)
  stack := []*File{rootFile}

  for len(stack) > 0 {
    file := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    children, _ := ioutil.ReadDir(file.Path)
    for _, child := range children {
      child := toFile(child, filepath.Join(file.Path, child.Name()))
      file.Children = append(file.Children, child)
      stack = append(stack, child)
    }
  }

	return *rootFile
}

func toFile(file os.FileInfo, path string) *File {
  JSONFile := File{ModifiedTime: file.ModTime(),
    IsDir:    file.IsDir(),
    Size:     file.Size(),
    Name:     file.Name(),
    Path:     path,
    Children: []*File{},
  }

  if file.Mode()&os.ModeSymlink == os.ModeSymlink {
    JSONFile.IsLink = true
    JSONFile.LinksTo, _ = filepath.EvalSymlinks(filepath.Join(path, file.Name()))
  }

  return &JSONFile
}
