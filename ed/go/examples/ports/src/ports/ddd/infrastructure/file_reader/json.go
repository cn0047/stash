package file_reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"ports/ddd/domain/model"
)

var (
	portRegExp = regexp.MustCompile(`(?is)"(.*)"\s*:\s.*({.*})`)
)

// todo: discuss this implementation.
func Read(input io.Reader, cb func(port model.PortEntity) error) error {
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanRunes)

	jsonBytes := make([]byte, 0)

	s.Scan() // scan rune '{'.
	for s.Scan() {
		b := s.Bytes()
		jsonBytes = append(jsonBytes, b...)
		if string(b) == "}" {
			var err error
			s.Scan() // scan runes `,` or `}`.

			p, err := Parse(jsonBytes)
			if err != nil {
				return fmt.Errorf("failed to parse data, error: %#v", err)
			}

			err = cb(p)
			if err != nil {
				return fmt.Errorf("failed to run callback with data, error: %#v", err)
			}

			jsonBytes = make([]byte, 0)
		}
	}

	return nil
}

func Parse(jsonBytes []byte) (model.PortEntity, error) {
	p := model.PortEntity{}
	matches := portRegExp.FindAllSubmatch(jsonBytes, -1)
	if len(matches) == 0 || len(matches[0]) < 2 {
		return p, fmt.Errorf("failed to match pord during parse stage, go jsonBytes: %s", jsonBytes)
	}

	err := json.Unmarshal(matches[0][2], &p)
	if err != nil {
		return p, fmt.Errorf("failed to unmarshal parsed port, error: %s", err)
	}

	p.ID = string(matches[0][1])

	return p, nil
}
