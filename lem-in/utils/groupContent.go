package utils

import (
	"errors"
	"strconv"
	"strings"
)

type Content struct {
	Ants      int
	Links     []string
	Rooms     []string
	StartRoom string
	EndRoom   string
}

func GroupContent(s []string) (Content, error) {
	var content Content

	ants, err := strconv.Atoi(s[0])
	if err != nil {
		return content, errors.New("wrong format")
	}

	content.Ants = ants

	for i, v := range s[1:] {
		if strings.Contains(v, "-") {
			content.Links = append(content.Links, v)
			continue
		}
		if v == "##start" {
			content.StartRoom = s[i+1]
			continue
		}
		if v == "##end" {
			content.EndRoom = s[i+1]
			continue
		}
		romms := strings.Split(v, " ")
		if len(romms) != 3 {
			return content, errors.New("wrong format")
		}
		content.Rooms = append(content.Rooms, romms[0])

	}

	return content, nil
}
