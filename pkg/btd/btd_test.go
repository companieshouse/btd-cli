package btd

import (
	"errors"
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitLoadTagMapWithEmptyFilePath(t *testing.T) {
	Convey("Given an empty tag map file path", t, func() {
		path := ""

		Convey("When loading the tag map", func() {
			tagMap, err := LoadTagMap(path)

			Convey("The value should be nil", func() {
				So(tagMap, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, errors.New("path cannot be empty"))
			})
		})
	})
}

func TestUnitLoadTagMapWithNonExistentFile(t *testing.T) {
	Convey("Given a file path for a non-existent file", t, func() {
		path := "non-existent"

		Convey("When loading the tag map", func() {
			tagMap, err := LoadTagMap(path)

			Convey("The value should be nil", func() {
				So(tagMap, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("unable to read tag map file: %s", path))
			})
		})
	})
}

func TestUnitLoadTagMapWithEmptyFile(t *testing.T) {
	Convey("Given a file path for an empty file", t, func() {

		f, err := os.CreateTemp("", "tmp-*")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		path := f.Name()

		Convey("When loading the tag map", func() {
			tagMap, err := LoadTagMap(path)

			Convey("The value should be nil", func() {
				So(tagMap, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("no mappings in tag map file: %s", path))
			})
		})
	})
}

func TestUnitLoadTagMapWithValidFile(t *testing.T) {
	Convey("Given a file path for a valid tag map file", t, func() {

		f, err := os.Open("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		path := f.Name()

		Convey("When loading the tag map", func() {
			tagMap, err := LoadTagMap(path)

			Convey("The value should not be nil", func() {
				So(tagMap, ShouldNotBeNil)
			})

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestUnitGetTagNameWithEmptyID(t *testing.T) {
	Convey("Given a valid tag map and empty ID", t, func() {

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		id := ""

		Convey("When retrieving the tag name", func() {
			name, err := tagMap.GetTagName(id)

			Convey("The name should be empty", func() {
				So(name, ShouldBeEmpty)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("id cannot be empty"))
			})
		})
	})
}

func TestUnitGetTagNameWithUnknownID(t *testing.T) {
	Convey("Given a valid tag map and unknown ID", t, func() {

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		id := "9999"

		Convey("When retrieving the tag name", func() {
			name, err := tagMap.GetTagName(id)

			Convey("The name should be empty", func() {
				So(name, ShouldBeEmpty)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("unknown id: %s", id))
			})
		})
	})
}

func TestUnitGetTagNameWithValidID(t *testing.T) {
	Convey("Given a valid tag map and ID", t, func() {

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		id := "0001"

		Convey("When retrieving the tag name", func() {
			name, err := tagMap.GetTagName(id)

			Convey("The tag name should not be empty", func() {
				So(name, ShouldNotBeNil)
			})

			Convey("The tag name should be correct", func() {
				So(name, ShouldEqual, "one")
			})

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestUnitLoadedFromFile(t *testing.T) {
	Convey("Given a file path and tag map", t, func() {

		path := "testdata/tagmap.dat"

		tagMap, err := LoadTagMap(path)
		if err != nil {
			t.Fatal(err)
		}

		Convey("When retrieving the file path the tag map was loaded from", func() {
			path := tagMap.LoadedFromFile()

			Convey("The file path should not be empty", func() {
				So(path, ShouldNotBeNil)
			})

			Convey("The path should be correct", func() {
				So(path, ShouldEqual, path)
			})
		})
	})
}

func TestUnitParseTagDataWithEmptyDataString(t *testing.T) {
	Convey("Given a tag map and empty BTD data string", t, func() {

		btd := ""

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		Convey("When parsing the BTD data string", func() {
			tagData, err := tagMap.ParseTagData(btd)

			Convey("The tag data should be nil", func() {
				So(tagData, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, errors.New("data string cannot be empty"))
			})
		})
	})
}

func TestUnitParseTagDataWithTruncatedIDField(t *testing.T) {
	Convey("Given a tag map and BTD data string with truncated ID", t, func() {

		id, len, data := "0", "", ""
		btd := fmt.Sprintf("%s%s%s", id, len, data)

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		Convey("When parsing the BTD data string", func() {
			tagData, err := tagMap.ParseTagData(btd)

			Convey("The tag data should be nil", func() {
				So(tagData, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, errors.New("reached EOF before parsing ID field"))
			})
		})
	})
}

func TestUnitParseTagDataWithTruncatedLengthField(t *testing.T) {
	Convey("Given a tag map and BTD data string with truncated length", t, func() {

		id, len, data := "0001", "", ""
		btd := fmt.Sprintf("%s%s%s", id, len, data)

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		Convey("When parsing the BTD data string", func() {
			tagData, err := tagMap.ParseTagData(btd)

			Convey("The tag data should be nil", func() {
				So(tagData, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("reached EOF before parsing length field for tag with id: %s", id))
			})
		})
	})
}

func TestUnitParseTagDataWithTruncatedDataField(t *testing.T) {
	Convey("Given a tag map and BTD data string with truncated data", t, func() {

		id, len, data := "0001", "0010", ""
		btd := fmt.Sprintf("%s%s%s", id, len, data)

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		Convey("When parsing the BTD data string", func() {
			tagData, err := tagMap.ParseTagData(btd)

			Convey("The tag data should be nil", func() {
				So(tagData, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("reached EOF before parsing data field for tag with id: %s", id))
			})
		})
	})
}

func TestUnitParseTagDataWithUknownIDField(t *testing.T) {
	Convey("Given a tag map and BTD data string containing an unknown ID", t, func() {

		id := "9999"
		btd := fmt.Sprintf("%s0004abcd", id)

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		Convey("When parsing the BTD data string", func() {
			tagData, err := tagMap.ParseTagData(btd)

			Convey("The tag data should be nil", func() {
				So(tagData, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("unknown id: %s", id))
			})
		})
	})
}

func TestUnitParseTagDataWithNonNumericIDField(t *testing.T) {
	Convey("Given a tag map and BTD data string containing a non-numeric ID field", t, func() {

		id, len, data := "xxxx", "0004", "abcd"
		btd := fmt.Sprintf("%s%s%s", id, len, data)

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		Convey("When parsing the BTD data string", func() {
			tagData, err := tagMap.ParseTagData(btd)

			Convey("The tag data should be nil", func() {
				So(tagData, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("found non-numeric id field: %s", id))
			})
		})
	})
}

func TestUnitParseTagDataWithNonNumericLengthField(t *testing.T) {
	Convey("Given a tag map and BTD data string containing a non-numeric length field", t, func() {

		id, len, data := "0001", "xxxx", "abcd"
		btd := fmt.Sprintf("%s%s%s", id, len, data)

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		Convey("When parsing the BTD data string", func() {
			tagData, err := tagMap.ParseTagData(btd)

			Convey("The tag data should be nil", func() {
				So(tagData, ShouldBeNil)
			})

			Convey("The error should describe the problem", func() {
				So(err, ShouldEqual, fmt.Errorf("found non-numeric length field for tag with id %s: %s", id, len))
			})
		})
	})
}

func TestUnitParseTagDataWithValidDataString(t *testing.T) {
	Convey("Given a tag map and BTD data string containing valid data", t, func() {

		id, len, data := "0001", "0004", "abcd"
		btd := fmt.Sprintf("%s%s%s", id, len, data)

		tagMap, err := LoadTagMap("testdata/tagmap.dat")
		if err != nil {
			t.Fatal(err)
		}

		Convey("When parsing the BTD data string", func() {
			tagData, err := tagMap.ParseTagData(btd)

			Convey("The tag data should not be nil", func() {
				So(tagData, ShouldNotBeNil)
			})

			Convey("The tag data should be correct", func() {
				So(tagData[0][0], ShouldEqual, "0001")
				So(tagData[0][1], ShouldEqual, "one")
				So(tagData[0][2], ShouldEqual, "0004")
				So(tagData[0][3], ShouldEqual, "abcd")
			})

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
