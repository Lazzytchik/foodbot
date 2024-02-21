package app

import "strings"

type QueryParser interface {
	Parse(query string) Params
}

type DefaultParser struct {
}

func (p *DefaultParser) Parse(query string) Params {
	var params Params

	chunks := strings.Split(strings.TrimSpace(query), " ")
	for _, chunk := range chunks {
		if p.IsCommand(chunk) {
			params.Path = chunk
			continue
		}

		if p.IsFlag(chunk) {
			// TODO
			continue
		}

		params.Params = append(params.Params, chunk)
	}

	return params
}

func (p DefaultParser) IsCommand(chunk string) bool {
	return len(chunk) > 1 && chunk[0] == '/'
}

func (p DefaultParser) IsFlag(chunk string) bool {
	return len(chunk) > 2 && chunk[0] == '-' && chunk[1] == '-'
}
