package trend

import (	
	"fmt"	
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

type Node struct {
	Attrs map[string]string
}
type Edge struct {
	Attrs map[string]string
}
type Graph struct {
	Title string
	Node  map[string]*Node
	Edge  map[string]map[string]*Edge
}

func NewGraph(name string) *Graph {
	log.Debugf("%s", name)
	return &Graph{
		Title: name,
		Node:  make(map[string]*Node),
		Edge:  make(map[string]map[string]*Edge),
	}
}

func (g *Graph) AddNode(name string, N *Node) error {
	log.Debugf("%s", name)
	g.Node[name] = N
	return nil
}

func (g *Graph) AddEdge(from string, to string, E *Edge) error {
	if _, exists := g.Edge[from]; !exists {
		g.Edge[from] = make(map[string]*Edge)
		g.Edge[from][to] = E
	} else {
		g.Edge[from][to] = E
	}
	return nil
}

func (g *Graph) PrintGrapData() error {
	for key, val := range g.Node {
		log.Debugf("[%s][%s]", key, val)
	}
	for k1, v1 := range g.Edge {
		for k2, v2 := range v1 {
			log.Debugf("[%s]-->[%s] [%s]", k1, k2, g.GetString(v2.Attrs))
		}
	}
	return nil
}

func (g *Graph) GetList(a map[string]string) []string {
	str := []string{}
	for k, atr := range a {
		str = append(str, fmt.Sprintf("%s=%q", k, atr))
	}
	return str
}
func (g *Graph) GetString(a map[string]string) string { //[괄호안의 속성]
	return strings.Join(g.GetList(a), " ")
}
func (g *Graph) GetLines(a map[string]string) string {
	str := strings.Join(g.GetList(a), "; ")
	return fmt.Sprintf("%s; ", str)
}
