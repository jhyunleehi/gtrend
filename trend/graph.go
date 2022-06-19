package trend

import (	
	"os"	

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
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
	//Node  map[string]*Node
	Node   []opts.GraphNode
	//Edge  map[string]map[string]*Edge
	Link []opts.GraphLink
}

func NewGraph(name string) *Graph {
	log.Debugf("%s", name)
	return &Graph{
		Title: name,
		Node:   []opts.GraphNode{},
		Link:  []opts.GraphLink{},
	}
}

func (g *Graph) AddNode( N opts.GraphNode) error {	
	g.Node=append(g.Node, N)
	return nil
}
func (g *Graph) AddLink(L opts.GraphLink) error {	
	g.Link=append(g.Link, L)
	return nil
}

func (g *Graph) PrintGrapData() error {
	for key, val := range g.Node {
		log.Debugf("[%s][%s]", key, val.Name)
	}
	for key, value := range g.Link {		
		log.Debugf("[%d] [%s]-->[%s]", key, value.Source,value.Target)		
	}
	return nil
}


