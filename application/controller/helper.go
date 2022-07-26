package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func GetPathParameter(path, prefix string) (int, error) {
	regex := regexp.MustCompile(fmt.Sprintf(`%s/(\d+)/?`, prefix))
	match := regex.FindAllStringSubmatch(path, -1)
	if len(match) == 0 || len(match[0]) < 2 { // ex. [[/aaa/1, 1]]
		return 0, fmt.Errorf("unmatch path . path: %s prefix: %s", path, prefix)
	}
	id, err := strconv.Atoi(match[0][1])
	if err != nil {
		return 0, err
	}
	return id, nil
}

func JsonDecode[T comparable](r *http.Request, w http.ResponseWriter, params *T) *T {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return params
}

func JsonEncode[T comparable](buff *bytes.Buffer, v T) error {
	enc := json.NewEncoder(buff)
	return enc.Encode(v)
}
