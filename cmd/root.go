/*
Copyright © 2025 Samuel Martineau and Dan Chiem

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ChiemMartineau/resgen/loader"
	"github.com/ChiemMartineau/resgen/renderer"
	"github.com/ChiemMartineau/resgen/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "resgen [path]",
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		rootPath := "."

		if len(args) > 0 {
			rootPath = args[0]
		}

		rootPath, err := filepath.Abs(rootPath)
		if err != nil {
			log.Fatal(utils.SmartFormatError(err))
		}

		data, err := loader.LoadData(rootPath)
		if err != nil {
			log.Fatal(utils.SmartFormatError(err))
		}

		// json, err := json.MarshalIndent(data, "", "\t")
		// if err != nil {
		// 	log.Fatal(utils.SmartFormatError(err))
		// }
		// fmt.Printf("%s\n", json)

		renderer.Text("./templates/resume/text.tmpl", data)
	},
}

func init() {
	log.SetFlags(0)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
