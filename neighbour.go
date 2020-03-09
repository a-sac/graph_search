package main

import (

	"fmt"
	"bytes"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
	"log"
	"strconv"
	"encoding/base64"
	"regexp"
	"os/exec"
	//"sort"
    "github.com/wcharczuk/go-chart"
	"github.com/soniakeys/graph"
	"github.com/soniakeys/graph/dot"

)

type NI int32

type AdjacencyList [][]NI

type Undirected struct {
    AdjacencyList
}

func random(min, max int) int32 {

    return int32(rand.Intn(max - min) + min)
}

func reverse(numbers []float64) []float64 {

	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {

		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}

func connectedRandomGraph(g graph.Undirected, n int32, m map[int32][]int32) {
	for !g.IsConnected() {

		var rand1 = random(0, int(n))
		var rand2 = random(0, int(n))

		for (rand1 == rand2) {

			rand2 = random(0, int(n))
		}

		has, _, _ := g.HasEdge(graph.NI(rand1),graph.NI(rand2))
		
		if !has {
			
			m[rand1] = append(m[rand1], rand2)
			m[rand2] = append(m[rand2], rand1)
            /*listy[rand] = listy[rand] + 1
			listy[r] = listy[r] + 1*/
			
			g.AddEdge(graph.NI(rand1), graph.NI(rand2))
		}
	}
}

func print_image(body []byte) {
	body64 := base64.StdEncoding.EncodeToString(body)
	term := os.Getenv("TERM")
	matched, err := regexp.MatchString("screen-\\w+",term)

	if err != nil {
		log.Fatal(err)
	}

	length := len(body)
	var buf bytes.Buffer
	
	if matched {
		buf.WriteString("\033Ptmux;\033\033]")
	} else{
		buf.WriteString("\033]")
	}
	
	buf.WriteString("1337;File=")
	buf.WriteString("size=")
	buf.WriteString(strconv.Itoa(length))
	buf.WriteString(";inline=1")
	buf.WriteString(":")
	buf.WriteString(body64)
	buf.WriteString("")
	buf.WriteString("\n")
	fmt.Println(buf.String())
}

func draw_graph(g graph.Undirected){
	c := exec.Command("dot", "-Tsvg", "-o", "al0.svg")
	w, err := c.StdinPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Start()

	dot.Write(g, w, dot.GraphAttr("rankdir", "LR"), dot.Isolated(true))
	w.Close()
	c.Wait()
}

func draw_chart(listx []float64, listy []float64, name string){
	graph := chart.Chart{

		XAxis: chart.XAxis{

			Style: chart.StyleShow(), //enables / displays the x-axis
		},
		YAxis: chart.YAxis{

			Style: chart.StyleShow(), //enables / displays the y-axis
		},
		Series: []chart.Series{

			chart.ContinuousSeries{

				XValues: listx,
				YValues: listy,
			},
		},
	}
	
	buffer := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buffer)
	err := ioutil.WriteFile(name, buffer.Bytes(), 0644)

	if err != nil {
        panic(err)
    }
}

func failure(args []string){
	var list []int32
	var neighbours []int32

	l, _ := strconv.Atoi(args[0])
	n := int32(l)
	l, _ = strconv.Atoi(args[1])
	rep := int32(l)
	failure, _ := strconv.ParseFloat(args[2], 64)

	rand.Seed(time.Now().Unix())

	var m map[int32][]int32 = make(map[int32][]int32)
	var g graph.Undirected 
	var j int32

	for j < n{
		list = append(list, 0)
		g.AddEdge(graph.NI(j),graph.NI(j))
		
        j += 1
	}

	connectedRandomGraph(g, n, m)

	var reps []int32
	var percentage []float64
	j=0
	for j < rep {
		alcanced := 0.0
		rand := random(0, int(n))
		neighbours = append(neighbours,rand)
		var i int32
		i = 0

		for(len(neighbours) > 0){
			i += 1
			var next []int32
			
			for _, node := range neighbours{
				neigh := float64(len(m[node])) * failure
				nneigh := int(neigh)
				for z, neighbour := range m[node] {
					if(z < nneigh){
						if list[neighbour]==0 {
							next = append(next, neighbour)
							list[neighbour] = i
							alcanced += 1
						}
					}
				}
			}

			neighbours = nil
			
			for _, node := range next{
				neighbours = append(neighbours, node)
			}
		}

		var k int32
		list = nil
		for k < n{
			list = append(list, 0)
			k +=1
		}
		
		percentage = append(percentage, alcanced/float64(n))
		reps = append(reps, i-1)
		j += 1
	}

	var media float64
	for _, r := range reps{
		media += float64(r)
	}

	media = media/float64(rep)

	fmt.Println("ROUNDS' AVERAGE: ", media)

	for i, r := range percentage{
		fmt.Println("ROUND ", i, ": We visited ", r * 100, "% of all nodes")
	}
}

func visited(args []string){
	var listx []float64
	var listy []float64
	var list []int32
	var neighbours []int32

	l, _ := strconv.Atoi(args[1])
	n := int32(l)

	rand.Seed(time.Now().Unix())

	var m map[int32][]int32 = make(map[int32][]int32)
	var g graph.Undirected 
	var j int32

	for j < n{
		list = append(list, 0)
		g.AddEdge(graph.NI(j),graph.NI(j))
		
        j += 1
	}

	connectedRandomGraph(g, n, m)

	var max float64
	max= 1.0
	var k float64
	k=0
	rand := random(0, int(n))
	for k <= max {
		alcanced:=0.0
		neighbours = append(neighbours,rand)
		var i int32
		i = 0

		for(len(neighbours) > 0){
			i += 1
			var next []int32
			
			for _, node := range neighbours{
				neigh := float64(len(m[node])) * k
				nneigh := int(neigh)
				for z, neighbour := range m[node] {
					if(z < nneigh){
						if list[neighbour]==0 {
							next = append(next, neighbour)
							list[neighbour] = i
							alcanced += 1
						}
					}
				}
			}

			neighbours = nil
			
			for _, node := range next{
				neighbours = append(neighbours, node)
			}
		}

		j=0
		list = nil
		for j < n{
			list = append(list, 0)
			j +=1
		}
		listx = append(listx, k)
		listy = append(listy, (alcanced/float64(n))*100)

		k += 0.01
		if(k>0.99 && k<1){
			k=1
		}
	}

	fmt.Println("CHECK visited.PNG")

	draw_chart(listx, listy, "visited.PNG")
}

func robustness(args []string){
	var listx []float64
	var listy []float64
	var list []int32
	var neighbours []int32

	l, _ := strconv.Atoi(args[1])
	n := int32(l)
	l, _ = strconv.Atoi(args[2])
	number_edges := int32(l)
	l, _ = strconv.Atoi(args[3])
	times := int32(l)

	rand.Seed(time.Now().Unix())

	var m map[int32][]int32 = make(map[int32][]int32)
	var g graph.Undirected 
	var j int32

	for j < n{
		list = append(list, 0)
		g.AddEdge(graph.NI(j),graph.NI(j))
		
        j += 1
	}

	connectedRandomGraph(g, n, m)

	var k int32
	k=0
	
	rand := random(0, int(n))
	for k <= times {
		var time int32 
		time =0
		for time < number_edges{
			has := true
			var rand1 int32
			var rand2 int32

			for has {
				rand1 = random(0, int(n))
				rand2 = random(0, int(n))
		
				for (rand1 == rand2) {
		
					rand2 = random(0, int(n))
				}
		
				has, _, _ = g.HasEdge(graph.NI(rand1),graph.NI(rand2))
			}

			m[rand1] = append(m[rand1], rand2)
			m[rand2] = append(m[rand2], rand1)
				
			g.AddEdge(graph.NI(rand1), graph.NI(rand2))

			time += 1
		}

		neighbours = append(neighbours,rand)
		var i int32
		i = 0

		for(len(neighbours) > 0){
			i += 1
			var next []int32
			
			for _, node := range neighbours{
				for _, neighbour := range m[node] {
					if list[neighbour]==0 {
						next = append(next, neighbour)
						list[neighbour] = i
					}
				}
			}

			neighbours = nil
			
			for _, node := range next{
				neighbours = append(neighbours, node)
			}
		}

		j=0
		list = nil
		for j < n{
			list = append(list, 0)
			j +=1
		}
		listx = append(listx, float64(k))
		listy = append(listy, float64(i-1))
		k += 1
	}

	fmt.Println("CHECK robustness.PNG")

	draw_chart(listx, listy, "robustness.PNG")
}

func main() {
	args := os.Args[1:]

	if len(args)==3 {
		failure(args)
	}else{
		if len(args)==2 && args[0]=="-visited"{
			visited(args)
		}else{
			if len(args)==4 && args[0]=="-robustness"{
				robustness(args)
			} else{
				fmt.Println("INVALID ARGUMENT USE!")
				fmt.Println("go run neighbour.go [NUMBER OF NODES] [NUMBER OF REPS] [PERCENTAGE OF NEIGHBOURS TO VISIT]")
				fmt.Println("go run neighbour.go -visited [NUMBER OF NODES]")
				fmt.Println("go run neighbour.go -robustness [NUMBER OF NODES] [NUMBER OF EDGES TO ADD] [NUMBER OF TIMES WE ADD EDGES]")
				b, _ := ioutil.ReadFile("fire.png")
				print_image(b)
				return
			}
		} 
	}

}