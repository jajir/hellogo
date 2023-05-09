package readsvg

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

func ReadNumber(data string) (float64, string, error) {
	data = TrimFromBeggining(data)
	if len(data) == 0 {
		return 0.0, data, errors.New("can't read number from empty string")
	}

	ch := data[0]
	strOut := bytes.NewBufferString("")

	if ch == '-' || ch == '+' {
		strOut.WriteByte(ch)
		data = data[1:]
		ch = data[0]
	}

	for (ch >= '0' && ch <= '9') || ch == '.' {
		strOut.WriteByte(ch)
		data = data[1:]
		if len(data) == 0 {
			ch = 255
		} else {
			ch = data[0]
		}
	}

	if strOut.Len() == 0 {
		return 0.0, data, fmt.Errorf("can't read number from '%s' string", data)
	}
	ven, err := strconv.ParseFloat(strOut.String(), 64)
	if err != nil {
		return 0.0, data, fmt.Errorf("can't read number from '%s' string, because of %w", data, err)
	}

	return ven, data, nil
}

func IsNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsNumberOrSign(ch byte) bool {
	return (ch >= '0' && ch <= '9') || ch == '+' || ch == '-'
}

func TrimFromBeggining(data string) string {
	if len(data) == 0 {
		return data
	}
	ch := data[0]
	for ch == ' ' {
		data = data[1:]
		if len(data) == 0 {
			return data
		}
		ch = data[0]
	}
	return data
}
