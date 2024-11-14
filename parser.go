package gorange

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// ExtractRange parses number ranges from string, supporting both single numbers and range notation
//
// Examples:
//   - Ignore non-digit noise: "xxx1-3xxx" => [1,2,3]
//   - Handle extra separators: "1,,2,,,3" => [1,2,3]
//   - Handle extra dashes: "1----3" => [1,2,3]
//   - Handle extra spaces: "  1  -  3  " => [1,2,3]
//   - Handle empty input: "" => []
//   - Handle range order: "3-1" => [1,2,3]
//   - Handle Chinese comma: "1，3-5，7" => [1,3,4,5,7]
//   - Handle 单双周: "1-15单周" => [1,3,5,7,9,11,13,15]
//
// !!!! Negative numbers are not supported
//
// Returns:
//
//	[]int: Slice containing all extracted numbers
//	error: Parsing errors (currently always returns nil)
func ExtractRange(input string, processRange ...func(rangeStr string, begin, end int) []int) ([]int, error) {
	// Chinese:'，' to English:','
	input = strings.ReplaceAll(input, "，", ",")

	result := make([]int, 0, 10)
	segments := strings.Split(input, ",")

	for _, segment := range segments {
		// Remove spaces at both ends and remove dashes in a loop
		segment = strings.TrimSpace(segment)
		for strings.HasPrefix(segment, "-") || strings.HasSuffix(segment, "-") {
			segment = strings.TrimPrefix(segment, "-")
			segment = strings.TrimSuffix(segment, "-")
		}
		segment = strings.TrimSpace(segment)

		// range
		pattern := `(\d+)[^\d]*-[^\d]*(\d+)`
		reRange := regexp.MustCompile(pattern)
		if reRange.MatchString(segment) {
			matchWeekRange := reRange.FindStringSubmatch(segment)
			if len(matchWeekRange) != 3 {
				log.Printf("Unable to match week range, %s", segment)
				continue
			}
			start, err := strconv.Atoi(matchWeekRange[1])
			if err != nil {
				log.Printf("Unable to Atoi, format: %s, err: %v", matchWeekRange[1], err)
				continue
			}
			end, err := strconv.Atoi(matchWeekRange[2])
			if err != nil {
				log.Printf("Unable to Atoi, format: %s, err: %v", matchWeekRange[2], err)
				continue
			}
			start, end = min(start, end), max(start, end)

			var tempRange []int
			if len(processRange) > 0 {
				tempRange = processRange[0](segment, start, end)
			} else {
				tempRange = DefaultProcessRange(segment, start, end)
			}

			result = append(result, tempRange...)
			continue
		}

		// single
		if segment != "" {
			pattern := `[^\d]*(\d+)[^\d]*`
			matchWeekSingle := regexp.MustCompile(pattern).FindStringSubmatch(segment)
			if len(matchWeekSingle) != 2 {
				log.Printf("Unable to find num, %s", segment)
				continue
			}
			atoi, err := strconv.Atoi(matchWeekSingle[1])
			if err != nil {
				log.Printf("Unable to Atoi, format: %s, err: %v", matchWeekSingle[2], err)
				continue
			}
			result = append(result, atoi)
		}
	}
	return result, nil
}

func DefaultProcessRange(rangeStr string, begin, end int) []int {
	result := make([]int, 0, end-begin+1)
	for i := begin; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

// 单双周处理
// 1-3单 => 1,3
func SingleDoubleWeekProcess(rangeStr string, begin, end int) []int {
	evenOrOdd := 999
	switch {
	case strings.Contains(rangeStr, "单"):
		evenOrOdd = 1
	case strings.Contains(rangeStr, "双"):
		evenOrOdd = 0
	}

	if evenOrOdd == 999 {
		return DefaultProcessRange(rangeStr, begin, end)
	}

	result := make([]int, 0, (end-begin+1)/2)
	for i := begin; i <= end; i++ {
		if i%2 == evenOrOdd {
			result = append(result, i)
		}
	}
	return result
}
