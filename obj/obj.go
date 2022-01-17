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
}

func Parse(s string) (*Parser, error) {
	p := &Parser{DefaultGroup: shape.NewGroup()}

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

			if len(verts) == 3 {
				p.DefaultGroup.AddChild(shape.NewTriangle(
					p.Vertices[verts[0]], p.Vertices[verts[1]], p.Vertices[verts[2]]),
				)
			}
		default:
			p.LinesIgnored++
		}
	}

	return p, nil
}
