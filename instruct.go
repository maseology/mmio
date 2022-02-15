package mmio

import (
	"log"
	"strings"
)

type Instruct struct {
	Desc  string
	Param map[string][]string
}

func NewInstruct(filepath string) *Instruct {
	a, err := ReadTextLines(filepath)
	if err != nil {
		log.Fatal(err)
	}
	ins := Instruct{Param: make(map[string][]string)}
	for _, s := range a {
		ts := strings.TrimSpace(s)
		if len(ts) > 0 && ts[0:1] != "!" {
			fs := strings.Fields(ts)
			nam := strings.ToLower(fs[0])
			switch nam {
			case "desc":
				ins.Desc = strings.Join(fs[1:], " ")
			default:
				c := 0
				for _, ss := range fs {
					if ss[0:1] == "!" {
						break
					}
					c += 1
				}
				ins.Param[nam] = fs[1:c]
			}
		}
	}
	return &ins
}
