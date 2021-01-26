package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// handler merupakan contoh controller
func handler(w http.ResponseWriter, r *http.Request) {
	name, err := getName(r.URL)
	if err != nil {
		Error(w, err)
		return
	}
	JSON(w, http.StatusOK, successResponse{
		Message: fmt.Sprintf("Hallo %s!", name),
	})
	return
}

// getName merupakan contoh service
func getName(u *url.URL) (string, error) {
	segment := u.Path[1:]
	if segment == "" {
		return "", ErrBadParamInput
	} else if segment == "foo" {
		return "foo", nil
	} else if segment == "bar" {
		// contoh handlling error lewat custom
		// cara 1 => return "", NewProblem(codeNotFound, http.StatusNotFound, fmt.Sprintf("%s not found!", segment))
		// cara 2 => return "", ErrNotFound
		// cara 3 => return "", NewErrNotFound(fmt.Sprintf("%s not found!", segment))
		return "", NewErrNotFound(fmt.Sprintf("%s not found!", segment))
	} else {
		// contoh handlling error dari eksternal
		i, err := strconv.ParseInt(segment, 10, 0)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("user # %d!", i), nil
	}
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8090", nil)
}
