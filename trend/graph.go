package trend

import (
	"bytes"
	"fmt"
	"io/ioutil"
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

func (g *Graph) CreateHpe3par(stg *model.StorageSpec) error {

	d := &hpe3par.Driver{}
	if err := d.Setup(stg); err == nil {
		defer d.Unset()
	}
	g.group = []string{"PORT", "VLAN", "IP", "VHOST", "VOL", "CPG", "RAID", "DISK"}

	portlist, _ := d.Client.GetPortsSys()
	vhostlist, _ := d.Client.GetVhostSys()
	vollist, _ := d.Client.GetVolumeSys()
	vlunlist, _ := d.Client.GetVLunSys()
	cpglist, _ := d.Client.GetCpgSys()
	disklist, _ := d.Client.GetPhysicalDisk()
	//Group
	for _, list := range g.group {
		g.AddNode(list, list, &Node{Attrs: map[string]string{"label": list, "shape": "plaintext"}})
	}
	//Port
	for _, port := range *portlist {
		if port.Type != 8 {
			continue
		}
		id := fmt.Sprintf("%d:%d:%d", port.PortPos.Node, port.PortPos.Slot, port.PartnerPos.CardPort)
		g.AddNode("PORT", id, &Node{Attrs: map[string]string{"label": id, "shape": "box"}})
	}
	//VLAN
	for _, port := range *portlist {
		portid := fmt.Sprintf("%d:%d:%d", port.PortPos.Node, port.PortPos.Slot, port.PartnerPos.CardPort)
		for _, vlan := range port.ISCSIVLans {
			id := fmt.Sprintf("VLAN:%s-%d", portid, vlan.VlanTag)
			g.AddNode("VLAN", id, &Node{Attrs: map[string]string{"label": id, "shape": "box"}})
		}
	}
	//IP
	for _, port := range *portlist {
		for _, vlan := range port.ISCSIVLans {
			id := fmt.Sprintf("IP:%s", vlan.IPAddr)
			g.AddNode("IP", id, &Node{Attrs: map[string]string{"label": id, "shape": "box"}})
		}
	}
	//VHOST
	for _, vhost := range *vhostlist {
		id := fmt.Sprintf("VHOST:%s", vhost.Name)
		g.AddNode("VHOST", id, &Node{Attrs: map[string]string{"label": id, "shape": "box"}})
	}
	//VOL
	for _, vol := range *vollist {
		id := fmt.Sprintf("VOL:%s", vol.VolumeName)
		g.AddNode("VOL", id, &Node{Attrs: map[string]string{"label": id, "shape": "box"}})
	}
	//CPG
	for _, cpg := range *cpglist {
		id := fmt.Sprintf("CPG:%s", cpg.VolumePoolName)
		g.AddNode("CPG", id, &Node{Attrs: map[string]string{"label": id, "shape": "box"}})
	}
	//RAID
	for _, cpg := range *cpglist {
		var raidtype int
		switch cpg.SDGrowth.LDlayout.RAIDType {
		case 1:
			raidtype = 0
		case 2:
			raidtype = 1
		case 3:
			raidtype = 5
		case 4:
			raidtype = 6

		}
		id := fmt.Sprintf("RAID:%d", raidtype)
		g.AddNode("RAID", id, &Node{Attrs: map[string]string{"label": id, "shape": "box"}})
	}
	//Disk
	for _, disk := range *disklist {
		id := fmt.Sprintf("DISK%d:%d:%d:%d", disk.CageID, disk.CageSide, disk.Mag, disk.DiskPos)
		g.AddNode("DISK", id, &Node{Attrs: map[string]string{"label": id, "shape": "box"}})
	}
	//Port -> VLAN
	for _, port := range *portlist {
		from := fmt.Sprintf("%d:%d:%d", port.PortPos.Node, port.PortPos.Slot, port.PartnerPos.CardPort)
		for _, vlan := range port.ISCSIVLans {
			to := fmt.Sprintf("VLAN:%s-%d", from, vlan.VlanTag)
			g.AddEdge(from, to, &Edge{Attrs: map[string]string{"label": from}})
		}
	}
	//VLAN->IP
	for _, port := range *portlist {
		portid := fmt.Sprintf("%d:%d:%d", port.PortPos.Node, port.PortPos.Slot, port.PartnerPos.CardPort)
		for _, vlan := range port.ISCSIVLans {
			from := fmt.Sprintf("VLAN:%s-%d", portid, vlan.VlanTag)
			to := fmt.Sprintf("IP:%s", vlan.IPAddr)
			g.AddEdge(from, to, &Edge{Attrs: map[string]string{"label": from}})
		}
	}
	//IP->HOST
	for _, vhost := range *vhostlist {
		for _, iscsipath := range vhost.ISCSIPaths {
			hostport := fmt.Sprintf("%d:%d:%d", iscsipath.PortPos.Node, iscsipath.PortPos.Slot, iscsipath.PortPos.CardPort)
			for _, port := range *portlist {
				portid := fmt.Sprintf("%d:%d:%d", port.PortPos.Node, port.PortPos.Slot, port.PortPos.CardPort)
				if hostport == portid {
					for _, vlan := range port.ISCSIVLans {
						from := fmt.Sprintf("IP:%s", vlan.IPAddr)
						to := fmt.Sprintf("VHOST:%s", vhost.Name)
						g.AddEdge(from, to, &Edge{Attrs: map[string]string{"label": from}})
					}
				}
			}
		}
	}
	//VHOST->VOL(VLUN)
	for _, vlun := range *vlunlist {
		from := fmt.Sprintf("VHOST:%s", vlun.HostName)
		to := fmt.Sprintf("VOL:%s", vlun.VolumeName)
		g.AddEdge(from, to, &Edge{Attrs: map[string]string{"label": from}})
	}

	//VOL->CPG
	for _, vol := range *vollist {
		from := fmt.Sprintf("VOL:%s", vol.VolumeName)
		to := fmt.Sprintf("CPG:%s", vol.UserCPG)
		g.AddEdge(from, to, &Edge{Attrs: map[string]string{"label": from}})
	}
	//CPG->RAID
	raidmap := map[string]string{}
	for _, cpg := range *cpglist {
		var raidtype int
		switch cpg.SDGrowth.LDlayout.RAIDType {
		case 1:
			raidtype = 0
		case 2:
			raidtype = 1
		case 3:
			raidtype = 5
		case 4:
			raidtype = 6
		}
		str := fmt.Sprintf("RAID:%d", raidtype)
		raidmap[str] = str
		from := fmt.Sprintf("CPG:%s", cpg.VolumePoolName)
		to := fmt.Sprintf("RAID:%d", raidtype)
		g.AddEdge(from, to, &Edge{Attrs: map[string]string{}})
	}
	//RAID->DISK
	for _, disk := range *disklist {
		to := fmt.Sprintf("DISK%d:%d:%d:%d", disk.CageID, disk.CageSide, disk.Mag, disk.DiskPos)
		for _, raid := range raidmap {
			from := fmt.Sprintf("%s", raid)
			g.AddEdge(from, to, &Edge{Attrs: map[string]string{}})
		}
	}
	return nil
}

func main() {
	//targets := []string{"netapp", "hpe3par", "ceph"}
	targets := []string{"netapp", "hpe3par"}
	storagelist := []model.StorageSpec{}
	
	for _, stg := range storagelist {
		g := NewGraph(stg.StorageVendor)
		switch stg.StorageVendor {
		case "hpe3par":

			g.CreateHpe3par(&stg)

			var buf bytes.Buffer
			g.WriteDot(&buf)
			//WriteDot(os.Stdout, g)
			fmt.Printf("%s", buf.String())
			filename := fmt.Sprintf("%s_%d", stg.StorageVendor, stg.StorageId)
			writeErr := ioutil.WriteFile(filename+".gv", []byte(buf.String()), 0755)
			if writeErr != nil {
				log.Fatalf("%v\n", writeErr)
			}
			DotToImage(filename, "svg", []byte(buf.String()))
		}
	}
}
