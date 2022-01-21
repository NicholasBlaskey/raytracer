package obj

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/nicholasblaskey/raytracer/shape"
	"github.com/nicholasblaskey/raytracer/tuple"
)

type Parser struct {
	LinesIgnored int
	Vertices     []tuple.Tuple
	Normals      []tuple.Tuple
	DefaultGroup *shape.Group
	Groups       map[string]*shape.Group
	curGroup     *shape.Group
}

func FileToGroup(filePath string) (*shape.Group, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	p, err := Parse(string(b))
	if err != nil {
		return nil, err
	}

	return p.ToGroup(), nil
}

func FileToBoundingBox(filePath string) (*shape.BoundingBox, error) {
	g, err := FileToGroup(filePath)
	if err != nil {
		return nil, err
	}

	return shape.NewBoundingBox(g), nil
}

func Parse(s string) (*Parser, error) {
	p := &Parser{DefaultGroup: shape.NewGroup(),
		Groups: make(map[string]*shape.Group)}
	p.Groups["DefaultGroup"] = p.DefaultGroup
	p.curGroup = p.DefaultGroup

	for _, line := range strings.Split(s, "\n") {
		line = strings.Trim(line, "\n\t\r ")
		if line == "" {
			continue
		}

		switch line[0] {
		case 'v':
			var err error
			vert := tuple.Point(0.0, 0.0, 0.0)
			startIndex := 2 // To account for it being vn
			if line[1] == 'n' {
				startIndex++
				vert = tuple.Vector(0.0, 0.0, 0.0)
			}

			verts := strings.Split(strings.Trim(line[startIndex:], "\n\t\r "), " ")

			if vert[0], err = strconv.ParseFloat(verts[0], 64); err != nil {
				return nil, err
			}
			if vert[1], err = strconv.ParseFloat(verts[1], 64); err != nil {
				return nil, err
			}
			if vert[2], err = strconv.ParseFloat(verts[2], 64); err != nil {
				return nil, err
			}

			if line[1] == 'n' {
				p.Normals = append(p.Normals, vert)
			} else {
				p.Vertices = append(p.Vertices, vert)
			}
		case 'f':
			vertStrings := strings.Split(line[2:], " ")
			var verts []int
			var norms []int
			for _, vertString := range vertStrings {
				split := strings.Split(vertString, "/")

				v, err := strconv.Atoi(split[0])
				if err != nil {
					return nil, err
				}
				verts = append(verts, v-1)

				// Face with normals case.
				if len(split) > 1 {
					n, err := strconv.Atoi(split[len(split)-1])
					if err != nil {
						return nil, err
					}
					norms = append(norms, n-1)
				}
			}

			// Triangle case.
			if len(verts) == 3 {
				if norms != nil {
					p.curGroup.AddChild(shape.NewSmoothTriangle(
						p.Vertices[verts[0]], p.Vertices[verts[1]], p.Vertices[verts[2]],
						p.Normals[norms[0]], p.Normals[norms[1]], p.Normals[norms[2]],
					))
				} else {
					p.curGroup.AddChild(shape.NewTriangle(
						p.Vertices[verts[0]], p.Vertices[verts[1]], p.Vertices[verts[2]]),
					)
				}
				continue
			}

			// Polygon case (fan triangulazation).
			for i := 1; i < len(verts)-1; i++ {
				if norms != nil {
					p.curGroup.AddChild(shape.NewSmoothTriangle(
						p.Vertices[verts[0]], p.Vertices[verts[i]], p.Vertices[verts[i+1]],
						p.Normals[norms[0]], p.Normals[norms[i]], p.Normals[norms[i+1]],
					))
				} else {
					p.curGroup.AddChild(shape.NewTriangle(
						p.Vertices[verts[0]], p.Vertices[verts[i]], p.Vertices[verts[i+1]]))
				}
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
