package obj

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Parser struct {
	LinesIgnored int
	Vertices     []tuple.Tuple
	DefaultGroup *shape.Group
	Groups       map[string]*shape.Group
	curGroup     *shape.Group
}

func Parse(s string) (*Parser, error) {
	p := &Parser{DefaultGroup: shape.NewGroup(),
		Groups: make(map[string]*shape.Group)}
	p.Groups["DefaultGroup"] = p.DefaultGroup
	p.curGroup = p.DefaultGroup

	for _, line := range strings.Split(s, "\n") {
		fmt.Println(line)
		line = strings.Trim(line, "\n\t\r ")
		if line == "" {
			continue
		}

		switch line[0] {
		case 'v':
			var err error
			verts := strings.Split(line[2:], " ")
			vert := tuple.Point(0.0, 0.0, 0.0)

			if vert[0], err = strconv.ParseFloat(verts[0], 64); err != nil {
				return nil, err
			}
			if vert[1], err = strconv.ParseFloat(verts[1], 64); err != nil {
				return nil, err
			}
			if vert[2], err = strconv.ParseFloat(verts[2], 64); err != nil {
				return nil, err
			}
			p.Vertices = append(p.Vertices, vert)
		case 'f':
			vertStrings := strings.Split(line[2:], " ")
			var verts []int
			for _, vertString := range vertStrings {
				v, err := strconv.Atoi(vertString)
				if err != nil {
					return nil, err
				}
				verts = append(verts, v-1)
			}

			// Triangle case.
			if len(verts) == 3 {
				p.curGroup.AddChild(shape.NewTriangle(
					p.Vertices[verts[0]], p.Vertices[verts[1]], p.Vertices[verts[2]]),
				)
				continue
			}

			// Polygon case (fan triangulazation).
			for i := 1; i < len(verts)-1; i++ {
				p.curGroup.AddChild(shape.NewTriangle(
					p.Vertices[verts[0]], p.Vertices[verts[i]], p.Vertices[verts[i+1]]))
			}
		case 'g':
			groupName := strings.Trim(line[2:], " \t\r")
			p.curGroup = shape.NewGroup()
			p.Groups[groupName] = p.curGroup

		default:
			p.LinesIgnored++
		}
	}

	return p, nil
}

func (p *Parser) ToGroup() *shape.Group {
	g := shape.NewGroup()
	for _, group := range p.Groups {
		g.AddChild(group)

		/*
			// Simplying assuming groups dont contain groups in models for now!
			for _, c := range group.Children {
				g.AddChild(c)
			}
		*/
	}
	return g
}
