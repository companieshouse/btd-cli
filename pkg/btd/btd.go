/*
Copyright © 2023 Companies House (Crown Copyright)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package btd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TagMap interface {
	ParseTagData(data string) (TagData, error)
	GetTagName(id string) (string, error)
	LoadTagMap(path string) (*TagMap, error)
	LoadedFromFile() string
}

type tagMapData struct {
	mappings map[string]string
	path     string
}

func (t *tagMapData) ParseTagData(data string) (TagData, error) {

	if len(data) == 0 {
		return nil, errors.New("data string cannot be empty")
	}

	var tagData TagData
	r := strings.NewReader(data)

	for {
		tag, err := parseTag(r, t)
		if err != nil {
			return nil, err
		}

		tagData = append(tagData, tag)

		if r.Len() == 0 {
			break
		}
	}

	return tagData, nil
}

func (t *tagMapData) GetTagName(id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("id cannot be empty")
	}

	name, ok := t.mappings[id]
	if !ok {
		return "", fmt.Errorf("unknown id: %s", id)
	}
	return name, nil
}

type TagData [][]string

func (t *TagData) GetMaxDataLength() int {
	max_data_length := 0
	for _, value := range *t {
		if data_length := len(value[3]); data_length > max_data_length {
			max_data_length = data_length
		}
	}

	return max_data_length
}

func LoadTagMap(path string) (*tagMapData, error) {
	tagMap := &tagMapData{make(map[string]string), path}

	if len(path) == 0 {
		return nil, errors.New("path cannot be empty")
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read tag map file: %s", path)
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)
	s.Split(bufio.ScanLines)

	pattern := regexp.MustCompile(`\s*([0-9]+)\s+([A-Za-z_]+)`)

	for s.Scan() {
		matches := pattern.FindStringSubmatch(s.Text())

		if len(matches) == 3 {
			id, tag := matches[1], matches[2]
			tagMap.mappings[id] = tag
		}
	}

	if len(tagMap.mappings) == 0 {
		return nil, fmt.Errorf("no mappings in tag map file: %s", path)
	}

	return tagMap, nil
}

func (t *tagMapData) LoadedFromFile() string {
	return t.path
}

func parseTag(r *strings.Reader, t *tagMapData) ([]string, error) {
	id, err := parseData(r, 4)
	if err != nil {
		return nil, errors.New("reached EOF before parsing ID field")
	}

	if _, err := parseUIntValue(id); err != nil {
		return nil, fmt.Errorf("found non-numeric id field: %s", id)
	}

	length, err := parseData(r, 4)
	if err != nil {
		return nil, fmt.Errorf("reached EOF before parsing length field for tag with id: %s", id)
	}

	lengthUnit, err := parseUIntValue(length)
	if err != nil {
		return nil, fmt.Errorf("found non-numeric length field for tag with id %s: %s", id, length)
	}

	data, err := parseData(r, uint(lengthUnit))
	if err != nil {
		return nil, fmt.Errorf("reached EOF before parsing data field for tag with id: %s", id)
	}

	tag, err := t.GetTagName(id)
	if err != nil {
		return nil, err
	}

	return []string{
		id,
		tag,
		length,
		data,
	}, nil
}

func parseData(r *strings.Reader, length uint) (string, error) {
	data := make([]byte, length)
	n, err := r.Read(data)
	if err == io.EOF || n < int(length) {
		return "", fmt.Errorf("not enough data remaining to read bytes: %d", length)
	}

	return string(data[:n]), nil
}

func parseUIntValue(data string) (uint64, error) {
	num, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		return 0, errors.New("unable to parse UInt value")
	}

	return num, nil
}
