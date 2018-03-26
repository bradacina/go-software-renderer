package obj

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	X float64
	Y float64
	Z float64
}

type Face struct {
	VertexIndex        [3]int
	VertexTextureIndex [3]int
}

type Obj struct {
	Vertices        []*Vertex
	Faces           []*Face
	VerticesTexture []*Vertex
}

func Load(filename string) *Obj {
	result := &Obj{}

	f, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " 	")
		line = strings.Replace(line, "\t", " ", -1)
		if len(line) == 0 {
			continue
		}

		if line[0] == 'v' {
			if line[1] == ' ' {
				parseVertex(line, result)
			}
			if line[1] == 't' {
				parseVertexTexture(line, result)
			}
		}

		if line[0] == 'f' {
			parseFace(line, result)
		}
	}

	return result
}

func parseVertexTexture(line string, obj *Obj) {
	line = strings.Replace(line, "  ", " ", -1)
	tokens := strings.Split(line, " ")
	if len(tokens) < 3 {
		log.Println("Encountered corrupt vertex texture at", line)
		return
	}

	x, err := strconv.ParseFloat(tokens[1], 32)
	if err != nil {
		log.Println(err, line)
		return
	}

	y, err := strconv.ParseFloat(tokens[2], 32)
	if err != nil {
		log.Println(err, line)
		return
	}

	obj.VerticesTexture = append(obj.VerticesTexture,
		&Vertex{X: float64(x), Y: float64(y)})
}

func parseVertex(line string, obj *Obj) {
	tokens := strings.Split(line, " ")
	if len(tokens) != 4 {
		log.Println("Encountered corrupt vertex at", line)
		return
	}

	x, err := strconv.ParseFloat(tokens[1], 32)
	if err != nil {
		log.Println(err, line)
		return
	}

	y, err := strconv.ParseFloat(tokens[2], 32)
	if err != nil {
		log.Println(err, line)
		return
	}

	z, err := strconv.ParseFloat(tokens[3], 32)
	if err != nil {
		log.Println(err, line)
		return
	}

	obj.Vertices = append(obj.Vertices, &Vertex{X: float64(x), Y: float64(y), Z: float64(z)})
}

func parseFace(line string, obj *Obj) {
	line = strings.Replace(line, "/", " ", -1)
	tokens := strings.Split(line, " ")
	if len(tokens) != 10 {
		log.Println("Encountered corrupt face at ", line)
		return
	}

	v1, err := strconv.Atoi(tokens[1])
	if err != nil {
		log.Println(err, line)
		return
	}

	vt1, err := strconv.Atoi(tokens[2])
	if err != nil {
		log.Println(err, line)
		return
	}

	v2, err := strconv.Atoi(tokens[4])
	if err != nil {
		log.Println(err, line)
		return
	}

	vt2, err := strconv.Atoi(tokens[5])
	if err != nil {
		log.Println(err, line)
		return
	}

	v3, err := strconv.Atoi(tokens[7])
	if err != nil {
		log.Println(err, line)
		return
	}

	vt3, err := strconv.Atoi(tokens[8])
	if err != nil {
		log.Println(err, line)
		return
	}

	obj.Faces = append(obj.Faces,
		&Face{VertexIndex: [3]int{v1, v2, v3}, VertexTextureIndex: [3]int{vt1, vt2, vt3}})
}
