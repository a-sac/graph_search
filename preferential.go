package main

import (

	//"fmt"
	"bytes"
	"io/ioutil"
	"math/rand"
	"time"
	"sort"
    "github.com/wcharczuk/go-chart"
	"github.com/soniakeys/graph"
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

func main() {

	var listx []float64
	var listy []float64
	var lst []int32

	rand.Seed(time.Now().Unix())

	var g graph.Undirected 
	n := int32(10000)
	var j int32

	for j < n{

		g.AddEdge(graph.NI(j),graph.NI(j))

		lst = append(lst, j)
        listx = append(listx, float64(j))
		listy = append(listy, float64(0))
		
        j += 1
	}

	for !g.IsConnected() {

		rand := random(0, int(n))
		
		l := len(lst)
		var r = lst[random(int(0), int(l))]

		for (rand == r) {

			r = lst[random(0, int(l))]
		}

		has, _, _ := g.HasEdge(graph.NI(r),graph.NI(rand))
		
		if !has && rand < n && r < n {

            listy[rand] = listy[rand] + 1
			listy[r] = listy[r] + 1
			
            lst = append(lst, rand)
			lst = append(lst, r)
			
			g.AddEdge(graph.NI(rand), graph.NI(r))
		}
	}

	sort.Float64s(listy)
	listy = reverse(listy)

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
	err := ioutil.WriteFile("chart.PNG", buffer.Bytes(), 0644)

	if err != nil {
        panic(err)
    }

}