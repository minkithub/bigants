package assertion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/kr/pretty"
)

var ErrFailed = fmt.Errorf("assertion: Failed")

func JSONDeepEqual(
	got interface{},
	expected interface{},
) (err error) {
	var (
		flatGot      interface{}
		flatExpected interface{}
		b            []byte
	)

	if b, err = json.Marshal(expected); err != nil {
		return err
	}
	if err = json.Unmarshal(b, &flatExpected); err != nil {
		return err
	}
	if b, err = json.Marshal(got); err != nil {
		return err
	}
	if err = json.Unmarshal(b, &flatGot); err != nil {
		return err
	}

	if !reflect.DeepEqual(flatGot, flatExpected) {
		var buff bytes.Buffer
		err := json.Indent(&buff, b, "", "  ")
		if err != nil {
			return err
		}
		desc := pretty.Diff(flatGot, flatExpected)
		return fmt.Errorf("%w %s\n%s",
			ErrFailed,
			string(buff.Bytes()),
			strings.Join(desc, "\n"),
		)
	}
	return nil
}
