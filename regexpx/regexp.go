// Package regexpx provides error-returning wrappers around regexp operations.
// Unlike the standard library's MustCompile functions, ordinary functions here
// return an error instead of panicking on invalid patterns.
package regexpx

import "regexp"

// Find compiles pattern and returns all submatches of the first match in text.
// Returns (nil, error) for an invalid pattern.
func Find(pattern string, text string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	match := re.FindStringSubmatch(text)
	if match == nil {
		return []string{}, nil
	}
	return match, nil
}

// FindString returns the first full match of pattern in text.
// Returns ("", false, nil) when there is no match.
func FindString(pattern string, text string) (string, bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", false, err
	}
	m := re.FindString(text)
	if m == "" {
		return "", false, nil
	}
	return m, true, nil
}

// FindAll compiles pattern and returns all matches with their subgroups.
// Returns (nil, error) for an invalid pattern.
func FindAll(pattern string, text string) ([][]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	matches := re.FindAllStringSubmatch(text, -1)
	if matches == nil {
		return [][]string{}, nil
	}
	return matches, nil
}

// Match reports whether text contains any match for pattern.
// Returns (false, error) for an invalid pattern.
func Match(pattern string, text string) (bool, error) {
	return regexp.MatchString(pattern, text)
}

// MustFind is like Find but panics if pattern is invalid.
func MustFind(pattern string, text string) []string {
	result, err := Find(pattern, text)
	if err != nil {
		panic(err)
	}
	return result
}
