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
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

type LocalizedString map[string]string

func (s *LocalizedString) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		// *s = []string{value.Value}
	case yaml.SequenceNode:
		// var list []string
		// for _, node := range value.Content {
		// 	list = append(list, node.Value)
		// }
		// *s = list
	default:
		return fmt.Errorf("unsupported YAML type for LocalizedString")
	}
	return nil
}

var rootCmd = &cobra.Command{
	Use:          "resgen [path]",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		source, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := goldmark.Convert(source, &buf); err != nil {
			return err
		}

		fmt.Println(buf.String())

		return nil
	},
}

func init() {

}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
