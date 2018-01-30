package command

import "fmt"

// splitState represents state
// for parsing with the Split() function
type splitState int

const (
	// outside of an argument
	splitStateNone splitState = 0
	// inside a quoted argument
	splitStateQuoted splitState = 1
	// inside a regular, unquoted argument
	splitStateUnQuoted splitState = 2
)

const (
	splitEscape      byte = '\\' // an escape character
	splitSingleQuote byte = '\'' // a single quote
	splitDoubleQuote byte = '"'  // a double quote

	splitSpace byte = ' '  // a space character
	splitTab   byte = '\t' // a tab character
)

// Split splits it into a list of arguments to be passed to os.exec
// based on https://stackoverflow.com/a/46973603
func Split(command string) ([]string, error) {
	var args []string // the arguments we will return
	current := ""     // the current argument

	state := splitStateNone // we start with no state at all
	stateQuote := byte(0)   // the quote of the current argument that was being used (if any)
	stateEscape := false    // flag indicating if the next character should be escaped

	// iterate over the characters in command
	for i := 0; i < len(command); i++ {
		c := command[i]

		// if the escape flag was set
		// the current character should be escaped
		if stateEscape {
			current += string(c)
			stateEscape = false
			continue
		}

		// if we get an escape character
		// the next character should be escaped
		// no matter the current state
		if c == splitEscape {
			stateEscape = true
			continue
		}

		// if we are in the quoted state
		// append the current character
		// unless we find the last quote
		if state == splitStateQuoted {
			if c != stateQuote {
				current += string(c)
			} else {
				args = append(args, current)
				current = ""
				state = splitStateNone
			}
			continue
		}

		// if we have a single or double quote
		// we should switch into the quote string
		if c == splitSingleQuote || c == splitDoubleQuote {
			state = splitStateQuoted
			stateQuote = c
			continue
		}

		// if we are in the argument state
		// we should end the current argument upon encountering spaces
		if state == splitStateUnQuoted {
			if c == splitSpace || c == splitTab {
				args = append(args, current)
				current = ""
				state = splitStateNone
			} else {
				current += string(c)
			}
			continue
		}

		// otherwise, we need to ignore everything that isn't space
		if c != splitSpace && c != splitTab {
			state = splitStateUnQuoted
			current += string(c)
		}
	}

	// we may not be inside a quoted argument
	if state == splitStateQuoted {
		return nil, fmt.Errorf("Unclosed quote in command line: %q", command)
	}

	// we may not ask for the next character to be escaped
	if stateEscape {
		return nil, fmt.Errorf("String missing escaped character: %q", command)
	}

	// add the last argument to the arguments
	if current != "" {
		args = append(args, current)
	}

	return args, nil
}
