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
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TagMap struct {
	mappings map[string]string
	path     string
}

func (t *TagMap) ParseTagData(data string) (TagData, error) {
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

func (t *TagMap) GetTagName(id string) (string, error) {
	name, ok := t.mappings[id]
	if !ok {
		return "", fmt.Errorf("no tag for id: %s", id)
	}
	return name, nil
}

type TagData [][]string

func LoadTagMap(path string) (*TagMap, error) {

	tagMap := &TagMap{make(map[string]string), path}

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

func (t *TagMap) LoadedFromFile() string {
	return t.path
}

func parseTag(r *strings.Reader, t *TagMap) ([]string, error) {
	id, err := parseData(r, 4)
	if err != nil {
		return nil, err
	}

	length, err := parseData(r, 4)
	if err != nil {
		return nil, err
	}

	length_uint, err := strconv.ParseUint(string(length), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("unable to parse length value for tag with id: %s", id)
	}

	data, err := parseData(r, uint(length_uint))
	if err != nil {
		return nil, err
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
	if err == io.EOF {
		return "", fmt.Errorf("reached EOF before parsing data")
	}

	return string(data[:n]), nil
}
