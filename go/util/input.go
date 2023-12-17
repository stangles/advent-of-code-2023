package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetIntInput(filename string) ([]int, error) {
	return GetIntInputWithSplitFunc(filename, bufio.ScanLines)
}

func GetIntInputWithSplitFunc(filename string, splitFunc bufio.SplitFunc) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file '%s': %w", filename, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file '%s': %v", filename, err)
		}
	}()

	lines := make([]int, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		text := scanner.Text()
		next, err := strconv.Atoi(text)
		if err != nil {
			return nil, fmt.Errorf("unable to convert value %v to int: %w", text, err)
		}

		lines = append(lines, next)
	}

	return lines, nil
}

func MustInt(ints []int, err error) []int {
	if err != nil {
		panic(err)
	}
	return ints
}

func GetStringInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file '%s': %w", filename, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file '%s': %v", filename, err)
		}
	}()

	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("unable to read contents of file '%s': %w", filename, err)
	}

	lines := strings.Split(string(contents), "\n")
	if lines[len(lines)-1] == "" {
		return lines[:len(lines)-1], nil
	} else {
		return lines, nil
	}
}

func MustString(strings []string, err error) []string {
	if err != nil {
		panic(err)
	}
	return strings
}

func GetBoards(filename string, boardSize int) ([][][]int, error) {
	input, err := GetStringInput(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to get input from '%s': %w", filename, err)
	}

	boardIdx := 0
	boards := make([][][]int, 0)
	boards = append(boards, make([][]int, boardSize))
	for pos, line := range input {
		if line == "" {
			boards = append(boards, make([][]int, boardSize))
			boardIdx++
			continue
		}

		row := pos
		if pos >= boardSize {
			row = (pos - boardIdx) % boardSize
		}
		boards[boardIdx][row] = make([]int, boardSize)
		for col, str := range strings.Fields(line) {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("unable to convert value in '%s': %w", filename, err)
			}

			boards[boardIdx][row][col] = num
		}
	}
	return boards, nil
}

func MustBoard(boards [][][]int, err error) [][][]int {
	if err != nil {
		panic(err)
	}
	return boards
}

func BinaryStrToUint16(binaryStrings []string) ([]uint16, error) {
	nums := make([]uint16, 0)
	for _, bin := range binaryStrings {
		num, err := strconv.ParseUint(bin, 2, 16)
		if err != nil {
			return nil, fmt.Errorf("unable to convert %s to number: %w", bin, err)
		}

		nums = append(nums, uint16(num))
	}

	return nums, nil
}

func MustBinaryStrToUint16(nums []uint16, err error) []uint16 {
	if err != nil {
		panic(err)
	}
	return nums
}

func ScanCommaSeparated(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		if len(data) > 0 && data[len(data)-1] == '\n' {
			return len(data), data[0 : len(data)-1], nil
		}
	}
	return 0, nil, nil
}
