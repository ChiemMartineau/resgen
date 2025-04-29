package utils

import (
	"fmt"
	"os"

	"github.com/ChiemMartineau/resgen/loader"
	"github.com/goccy/go-yaml/printer"
	"golang.org/x/term"
)

func SmartFormatError(err error) string {
	if term.IsTerminal(int(os.Stderr.Fd())) {
		return PrettyFormatError(err)
	} else {
		return err.Error()
	}
}

func PrettyFormatError(err error) string {
	switch err := err.(type) {
	case loader.YamlError:
		var pp printer.Printer
		src := pp.PrintErrorToken(err.YamlError.GetToken(), true)
		return fmt.Sprintf("%s%s%s\n%s", ColourRedBold, err.Error(), ColourReset, src)
	default:
		return fmt.Sprintf("%s%s%s", ColourRedBold, err.Error(), ColourReset)
	}
}
