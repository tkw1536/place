package utils

import (
	"log"
	"os"
)

// Logger is a single Logger that is used by all place packages
var Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
