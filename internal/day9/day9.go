package day9

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/filereader"
)

func Main() {
	log.Println("DAY 9")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput(9, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}

	dm := lines[0]

	disk := []int{}
	for i, v := range dm {
		n, err := strconv.Atoi(string(v))

		if err != nil {
			log.Fatalln(err)
		}
		if i%2 == 0 {
			id := i / 2
			for range n {
				disk = append(disk, id)
			}
		} else {
			for range n {
				disk = append(disk, -1)
			}
		}
	}

	start := 0
	end := len(disk) - 1
	for disk[end] == -1 {
		end--
	}
	for start < end {
		if disk[start] == -1 {
			disk[start] = disk[end]
			disk[end] = -1
			end--
			for disk[end] == -1 {
				end--
			}
		}
		start++
	}

	solution := 0
	for i, id := range disk {
		if id != -1 {
			solution += id * i
		}
	}
	log.Println("The solution is:", solution)
}

type block struct {
	id   int
	size int
}

func printDisk(blocks []block) {
	for _, block := range blocks {
		if block.id == -1 {
			fmt.Print(strings.Repeat(".", block.size))
		} else {
			fmt.Print(strings.Repeat(fmt.Sprintf("%d", block.id), block.size))
		}
	}
	fmt.Println()
}

func part2() {
	f := filereader.NewFromDayInput(9, 1)
	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}

	dm := lines[0]
	disk := []int{}
	for i, v := range dm {
		n, err := strconv.Atoi(string(v))

		if err != nil {
			log.Fatalln(err)
		}
		if i%2 == 0 {
			id := i / 2
			for range n {
				disk = append(disk, id)
			}
		} else {
			for range n {
				disk = append(disk, -1)
			}
		}
	}

	blocks := []block{}
	id := disk[0]
	n := 0
	for _, v := range disk {
		if v != id {
			blocks = append(blocks, block{id: id, size: n})
			n = 0
			id = v
		}
		n++
	}
	blocks = append(blocks, block{id: id, size: n})

	for end := len(blocks) - 1; end >= 0; end-- {
		if blocks[end].id == -1 {
			continue
		}
		for start := range end {
			if blocks[start].id != -1 {
				continue
			}
			if blocks[end].size <= blocks[start].size {
				blockEnd := blocks[end]
				blocks[end].id = -1
				if end+1 < len(blocks) && blocks[end+1].id == -1 {
					blocks[end].size += blocks[end+1].size
					blocks = slices.Delete(blocks, end+1, end+2)
				}
				if blocks[end-1].id == -1 {
					blocks[end-1].size += blocks[end].size
					blocks = slices.Delete(blocks, end, end+1)
                    end--
				}
				blocks[start].size -= blockEnd.size
				if blocks[start].size == 0 {
					blocks = slices.Delete(blocks, start, start+1)
					end--
				}
				blocks = slices.Insert(blocks, start, blockEnd)
                end++
				break
			}
		}
	}

	solution := 0
	index := 0
	for _, block := range blocks {
		if block.id == -1 {
			index += block.size
		} else {
			for range block.size {
				solution += block.id * index
				index++
			}
		}
	}
	log.Println("The solution is:", solution)
}
