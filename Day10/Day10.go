package main 

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(file))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(file))
}

func Part1Solver(pipes []string) int {
	var start []int = getPos(pipes, 'S');
	direction := startMoves(pipes, start);
	moves := [2][]int{};

	for i := 0; i < 2; i++ {
		for _, num := range start {
			moves[i] = append(moves[i], num);
		}
		moves[i], direction[i] = move(moves[i], pipes, direction[i]);
	}
	
	path := 1;

	for ; moves[0][0] != moves[1][0] || moves[0][1] != moves[1][1]; path++ {
		moves[0], direction[0] = move(moves[0], pipes, direction[0]);
		if moves[0][0] == moves[1][0] && moves[0][1] == moves[1][1] {
			break;
		}
		moves[1], direction[1] = move(moves[1], pipes, direction[1]);
	}
	return path;
}

func Part2Solver(pipes []string) int {
	var boundedArea = make([][]int, len(pipes));
	for j := 0; j < len(pipes); j++ {
		temp := make([]int, len(pipes[j]));
		boundedArea[j] = temp;
	}
	var start []int = getPos(pipes, 'S');
	
	boundedArea[start[1]][start[0]] = 1;	

	direction := startMoves(pipes, start);
	moves := [2][]int{};

	for i := 0; i < 2; i++ {
		for _, num := range start {
			moves[i] = append(moves[i], num);
		}
		moves[i], direction[i] = move(moves[i], pipes, direction[i]);
		boundedArea[moves[i][1]][moves[i][0]] = 1;
	}
	

	for path := 0; moves[0][0] != moves[1][0] || moves[0][1] != moves[1][1]; path++ {
		moves[0], direction[0] = move(moves[0], pipes, direction[0]);
		boundedArea[moves[0][1]][moves[0][0]] = 1;
		if moves[0][0] == moves[1][0] && moves[0][1] == moves[1][1] {
			break;
		}
		moves[1], direction[1] = move(moves[1], pipes, direction[1]);
		boundedArea[moves[1][1]][moves[1][0]] = 1;
	}

	boundedA := acrossBounds(pipes, boundedArea, startMoves(pipes, start));
	boundedU := downBounds(pipes, boundedArea, startMoves(pipes, start));

	if boundedA != boundedU {
		log.Println("something went wrong, the vertical and horizontal interiors should be the same");
	}
	return boundedA;
}

func acrossBounds(grid []string, bounds [][]int, startType []byte) int {
	startMove := [2]int{0, 0};
	for _, move := range startType {
		if move == 'U' {
			startMove[0] = 1;
		} else if move == 'D' {
			startMove[1] = 1;
		}
	}
	interior := 0;
	for j := 0; j < len(grid); j++ {
		enclosed := [2]int{0, 0};
		for i := 0; i < len(grid[j]); i++ {
			if bounds[j][i] == 1 {
				switch grid[j][i] {
					case '|':
						enclosed[0]++;
						enclosed[1]++;
					case 'F':
						enclosed[1]++;
					case '7':
						enclosed[1]++;
					case 'L':
						enclosed[0]++;
					case 'J':
						enclosed[0]++;
					case 'S':
						enclosed[0] += startMove[0];
						enclosed[1] += startMove[1];
					}
			} else if enclosed[0] % 2 != 0 || enclosed[0] % 2 != 0 { 
				interior++; 
			}
		}
	}
	return interior;
}

func downBounds(grid []string, bounds [][]int, startType []byte) int {
	startMove := [2]int{0, 0};
	for _, move := range startType {
		if move == 'L' {
			startMove[0] = 1;
		} else if move == 'R' {
			startMove[1] = 1;
		}
	}
	interior := 0;
	for i := 0; i < len(grid[0]); i++ {
		enclosed := [2]int{0, 0};
		for j := 0; j < len(grid); j++ {
			if bounds[j][i] == 1 {
				switch grid[j][i] {
					case '-':
						enclosed[0]++;
						enclosed[1]++;
					case 'F':
						enclosed[1]++;
					case '7':
						enclosed[0]++;
					case 'L':
						enclosed[1]++;
					case 'J':
						enclosed[0]++;
					case 'S':
						enclosed[0] += startMove[0];
						enclosed[1] += startMove[1];
					}
			} else if enclosed[0] % 2 != 0 || enclosed[0] % 2 != 0 { 
				interior++; 
			}
		}
	}
	return interior;
}

func move(pos []int, grid []string, move byte) ([]int, byte) {
	switch move {
		case 'U':
			pos[1] -= 1;
		case 'D':
			pos[1] += 1;
		case 'L': 
			pos[0] -= 1; 
		case 'R': 
			pos[0] += 1;
		}
	var nextMove byte = '|';
	switch grid[pos[1]][pos[0]] {
		case '|':
			nextMove = move;	
		case '-':
			nextMove = move;	
		case 'L':
			if move == 'D' {
				nextMove = 'R';
			} else {
				nextMove = 'U';
			}
		case 'J':
			if move == 'D' {
				nextMove = 'L';
			} else {
				nextMove = 'U';
			}
		case '7':
			if move == 'U' {
				nextMove = 'L';
			} else {
				nextMove = 'D';
			}
		case 'F':
			if move == 'U' {
				nextMove = 'R';
			} else {
				nextMove = 'D';
			}
		default:
			fmt.Println(grid[pos[1]][pos[0]]);
			fmt.Println(string(grid[pos[1]][pos[0]]));
	}
	return pos, nextMove;
}

func startMoves(pipes []string, start []int) []byte {
	pipeDir := []byte{};
	if start[0] > 0 && contains("-FL", pipes[start[1]][start[0] - 1]) {
		pipeDir = append(pipeDir, 'L');
	}
	if start[1] > 0 && contains("|F7", pipes[start[1] - 1][start[0]]) {
		pipeDir = append(pipeDir, 'U');
	}
	if start[0] < len(pipes[0]) - 1 && contains("-7J", pipes[start[1]][start[0] + 1]) {
		pipeDir = append(pipeDir, 'R');
	}
	if start[1] < len(pipes) - 1 && contains("|LJ", pipes[start[1] + 1][start[0]]) {
		pipeDir = append(pipeDir, 'D');
	}
	return pipeDir;
}

func contains(str string, char byte) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			return true;
		}
	}
	return false;
}
			
func getPos(grid []string, char byte) []int {
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if grid[j][i] == char {
				return []int {i, j};
			}
		}
	}
	return []int {-1, -1};
}

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func readFile(name string) []string {
	file, err := os.Open(name + ".txt")
	if err != nil {
		log.Println("Failed to read file as:", err)
		panic(err)
	}
	defer file.Close()
	fileContent := bufio.NewScanner(file)
	content := []string{}
	for fileContent.Scan() {
		content = append(content, fileContent.Text())
	}
	return content
}
