/*
	gonav - A Source Engine navigation mesh file parser written in Go.
	Copyright (C) 2016  Matt Razza

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published
	by the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/mrazza/gonav"
)

func writeLines(lines [][]float32, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, curRow := range lines {
	    for _, line := range curRow {
	        fmt.Fprintln(w, line)
	    }
	}
    return w.Flush()
}


func main() {
	// /Users/adithyansujithkumar/Documents/Projects/fps/csgo_stuff/csgo/csgo/data/nav/de_vertigo.nav
	fmt.Print("Enter file name (ex: ../de_vertigo.nav): ")
	var file string
	fmt.Scanf("%s\n", &file)
	f, ok := os.Open(file) // Open the file

	if ok != nil {
		fmt.Printf("Failed to open file: %v\n", ok)
		return
	}

	defer f.Close()
	//start := time.Now()
	parser := gonav.Parser{Reader: f}
	mesh, _ := parser.Parse() // Parse the file
	point := gonav.Vector3{X: 0, Y: 0, Z: 0}
	area := mesh.GetNearestArea(point, true)
	if area != nil {
		//fmt.Printf("Found in %fus...\n", float64(elapsed.Nanoseconds())/10000.0)
		fmt.Println(area)
		fmt.Printf("----")
		tmpVal := 0
		maxVal := uint32(0)


		//h := &IntHeap{}
		//var map_1 map[int]int
		for _, curr := range mesh.Areas {
			//fmt.Println(curr.ID, maxVal)
			//heap.Push(h, curr.ID)
			if curr.ID > maxVal {
				maxVal = curr.ID
			}
			tmpVal++
			
		}
		

		fmt.Printf("----")
		fmt.Println(tmpVal, maxVal)
		distMat := make([][]float32, maxVal + 1) // One row per unit of y.
		// Loop over the rows, allocating the slice for each row.
		for i := range distMat {
			distMat[i] = make([]float32, maxVal + 1)
		}

		for _, outer := range mesh.Areas {
			//fmt.Println(outer.ID)
			for _, inner := range mesh.Areas {
				curPath, _ := gonav.SimpleBuildShortestPath(outer, inner)
				if len(curPath.Nodes) > 0 {
					distMat[outer.ID][inner.ID] = curPath.Nodes[len(curPath.Nodes)-1].CostFromStart
				} else {
					distMat[outer.ID][inner.ID] = 0.0
				}

				//fmt.Println(distMat[outer.ID][inner.ID])
			}
			
		}
		writeLines(distMat, "tmp.txt")
	

	} else {
		fmt.Printf("Nope")
	}
}