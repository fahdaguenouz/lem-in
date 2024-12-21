package Mosdef

import (
	"log"
	"strconv"
	"strings"
)

func GetRooms(lines []string) (string, string, int, map[string][]string) {
	antsnumber := 0
	startroom := false
	endroom := false
	start := ""
	end := ""
	rooms := make(map[string]struct{})
	coords := make(map[string]struct{})
	links := make(map[string][]string)
	for i, line := range lines {
		if len(strings.Split(line, " ")) != 3 && len(strings.Split(line, "-")) != 2 && line[0] != '#' && i != 0 {

			log.Fatal("ERROR: invalid data format. invalid syntax")
		}
		if i == 0 {
			number, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("ERROR: invalid data format. invalid ants number")
			}
			if number <= 0 {
				log.Fatal("ERROR: invalid data format. invalid number of ants")
			}
			antsnumber = number
		}
		if len(strings.Split(line, " ")) == 3 && line[0] != '#' {
			if _, exist := rooms[strings.Split(line, " ")[0]]; exist {
				log.Fatal("ERROR: invalid data format. room already exists")
			}
			rooms[strings.Split(line, " ")[0]] = struct{}{}
			if _, exist := coords[strings.Split(line, " ")[1]+"+"+strings.Split(line, " ")[2]]; exist {
				log.Fatal("ERROR: invalid data format. repeated coordinates\n" + line + "\n" + strconv.Itoa(i))
			}

			coords[strings.Split(line, " ")[1]+"+"+strings.Split(line, " ")[2]] = struct{}{}
		}
		if line == "##start" && !startroom {
			if i+1 != len(lines) {
				startroom = true
				start = strings.Split(lines[i+1], " ")[0]
			} else {
				log.Fatal("ERROR: invalid data format.there is no start room")
			}
		} else if line == "##start" && startroom {
			log.Fatal("ERROR: invalid data format.there is too many start rooms")
		}
		if line == "##end" && !endroom {
			if i+1 != len(lines) && len(strings.Split(lines[i+1], " ")) == 3 {
				endroom = true
				end = strings.Split(lines[i+1], " ")[0]
			} else {
				log.Fatal("ERROR: invalid data format.there is no end room")
			}
		} else if line == "##end" && endroom {
			log.Fatal("ERROR: invalid data format.there is too many end rooms")
		}

		if len(strings.Split(line, "-")) == 2 && line[0] != '#' {
			if _, exist := rooms[strings.Split(line, "-")[0]]; !exist {
				log.Fatal("ERROR: invalid data format. unexisting room")
			}
			if _, exist := rooms[strings.Split(line, "-")[1]]; !exist {
				log.Fatal("ERROR: invalid data format. unexisting room")
			}
			links[strings.Split(line, "-")[0]] = append(links[strings.Split(line, "-")[0]], strings.Split(line, "-")[1])
			links[strings.Split(line, "-")[1]] = append(links[strings.Split(line, "-")[1]], strings.Split(line, "-")[0])
		}
	}
	if len(links) != len(rooms) {
		log.Fatal("ERROR: invalid data format. room without links room")
	}

	return start, end, antsnumber, links
}
