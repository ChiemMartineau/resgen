/*
Copyright © 2025 Samuel Martineau <samuel.martineau@uwaterloo.ca> and Dan Chiem <dchiem@uwaterloo.ca>

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
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/spf13/cobra"
)

// genpdfCmd represents the genpdf command
var genpdfCmd = &cobra.Command{
	Use:   "genpdf",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := chromedp.NewContext(cmd.Context(), chromedp.WithLogf(log.Printf))
		defer cancel()

		var buf []byte
		if err := chromedp.Run(ctx, printToPDF("file:///home/samuelm/Desktop/resgen/cv.html", &buf)); err != nil {
			log.Fatal(err)
		}

		if err := os.WriteFile("sample.pdf", buf, 0o644); err != nil {
			log.Fatal(err)
		}

		fmt.Println("wrote sample.pdf")
	},
}

func init() {
	rootCmd.AddCommand(genpdfCmd)
}

func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithMarginTop(0).
				WithMarginRight(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
