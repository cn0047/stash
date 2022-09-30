package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// RegExpSubSetCommand captures single command enclosed in square brackets.
	RegExpSubSetCommand = regexp.MustCompile(`(?U)\[[^\[]+\]`)
)

func main() {
	commandLine := strings.Join(os.Args[1:], " ")

	for {
		m := findSubSetCommand(commandLine)
		if len(m) == 0 {
			break
		}
		subSetCommand := m[0][0] // @TBD: Intentionally brittle.
		resultFileName, err := handleCommand(subSetCommand)
		if err != nil {
			log.Fatalf("failed to print reslut, error: %v", err)
		}
		// Replace sub-command with filename which contains calculated result.
		commandLine = strings.Replace(commandLine, subSetCommand, resultFileName, -1)
	}

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	err := loadFromFile(commandLine, func(val int) {
		_, err := fmt.Fprintf(writer, "%d\n", val)
		if err != nil {
			log.Fatalf("failed to print value, error: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("failed to print result, error: %v", err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatalf("failed to flush writer, error: %v", err)
	}
}

// findSubSetCommand returns first found sub-command from command line.
func findSubSetCommand(s string) [][]string {
	m := RegExpSubSetCommand.FindAllStringSubmatch(s, -1)
	return m
}

// handleCommand performs command handling and returns filename which contains computed result.
func handleCommand(cmd string) (resultFileName string, err error) {
	s := strings.Split(cmd, " ")
	commandName := s[1]
	set1 := s[2]
	set2 := s[3]

	var resultSet []int
	switch commandName {
	case "INT":
		resultSet, err = intsc(set1, set2)
	case "SUM":
		resultSet, err = sum(set1, set2)
	case "DIF":
		resultSet, err = sdiff(set1, set2)
	default:
		return "", fmt.Errorf("got unsupported command: %v", commandName)
	}
	if err != nil {
		return "", fmt.Errorf("failed to fulfill command %s, error: %w", commandName, err)
	}

	resultFileName, err = saveToTmpFile(resultSet)
	if err != nil {
		return "", fmt.Errorf("failed to save resultset into file, error: %v", err)
	}

	return resultFileName, err
}

// saveToTmpFile saves slice of int values to temporary file and returns filename.
func saveToTmpFile(data []int) (string, error) {
	f, err := ioutil.TempFile(os.TempDir(), "scalc")
	if err != nil {
		return "", fmt.Errorf("failed to init tmp file, error: %w", err)
	}
	defer f.Close()

	for _, v := range data {
		_, err = f.WriteString(fmt.Sprintf("%d\n", v))
		if err != nil {
			return "", fmt.Errorf("failed to write to tmp file, error: %w", err)
		}
	}

	return f.Name(), nil
}

// loadIntoSlice loads data from file into int values slice.
func loadIntoSlice(fileName string) ([]int, error) {
	s := make([]int, 0) // @TBD: Optimization.

	err := loadFromFile(fileName, func(val int) {
		s = append(s, val)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load map from file, error: %w", err)
	}

	return s, nil
}

// loadIntoMap loads data from file into int values map.
func loadIntoMap(fileName string) (map[int]struct{}, error) {
	m := make(map[int]struct{}) // @TBD: Optimization.

	err := loadFromFile(fileName, func(val int) {
		m[val] = struct{}{}
	})
	if err != nil {
		return nil, fmt.Errorf("failed to load map from file, error: %w", err)
	}

	return m, nil
}

// loadFromFile loads data from file and call callback for each line from file.
func loadFromFile(fileName string, callback func(val int)) error {
	file, err := os.Open(fileName) // @TBD: Intentionally don't check file existence.
	if err != nil {
		return fmt.Errorf("failed to open file, error: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line := ""

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		cleanLine := strings.Replace(line, "\n", "", -1)
		val, err := strconv.Atoi(cleanLine)
		if err != nil {
			return fmt.Errorf("failed to parse int value from string: %v", cleanLine)
		}

		callback(val)
	}

	return nil
}

// intsc provides intersection for 2 sets.
func intsc(set1 string, set2 string) ([]int, error) {
	setA, err := loadIntoMap(set1)
	if err != nil {
		return nil, fmt.Errorf("failed to load set into map, error: %w", err)
	}

	setB, err := loadIntoSlice(set2)
	if err != nil {
		return nil, fmt.Errorf("failed to load set into slice, error: %w", err)
	}

	size := len(setA)
	if len(setB) > size {
		size = len(setB)
	}
	result := make([]int, 0, size)

	for i := 0; i < len(setB); i++ {
		val := setB[i]
		_, inSetA := setA[val]
		if inSetA {
			result = append(result, val)
		}
	}

	return result, nil
}

// diff provides difference for 2 sets.
func diff(set1 string, set2 string) ([]int, error) {
	setA, err := loadIntoSlice(set1)
	if err != nil {
		return nil, fmt.Errorf("failed to load set into slice, error: %w", err)
	}

	setB, err := loadIntoMap(set2)
	if err != nil {
		return nil, fmt.Errorf("failed to load set into map, error: %w", err)
	}

	result := make([]int, 0, len(setA))

	for i := 0; i < len(setA); i++ {
		val := setA[i]
		_, inSetB := setB[val]
		if !inSetB {
			result = append(result, val)
		}
	}

	return result, nil
}

// sdiff provides symmetric difference for 2 sets.
func sdiff(set1 string, set2 string) ([]int, error) {
	r1, err := diff(set1, set2)
	if err != nil {
		return nil, fmt.Errorf("failed to fulfill diff, error: %w", err)
	}

	r2, err := diff(set2, set1)
	if err != nil {
		return nil, fmt.Errorf("failed to fulfill diff, error: %w", err)
	}

	return append(r1, r2...), nil
}

// sum provides union for 2 sets.
func sum(set1 string, set2 string) ([]int, error) {
	setA, err := loadIntoMap(set1)
	if err != nil {
		return nil, fmt.Errorf("failed to load set into map, error: %w", err)
	}

	setB, err := loadIntoSlice(set2)
	if err != nil {
		return nil, fmt.Errorf("failed to load set into slice, error: %w", err)
	}

	result, err := loadIntoSlice(set1) // @TBD: Optimization.
	if err != nil {
		return nil, fmt.Errorf("failed to load set into slice, error: %w", err)
	}

	for i := 0; i < len(setB); i++ {
		val := setB[i]
		_, inSetA := setA[val]
		if !inSetA {
			result = append(result, val)
		}
	}

	return result, nil
}
