package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	arr := strings.Split(repeat, " ")
	if len(arr) == 0 {
		return "", nil
	}

	typeOfRepeat := arr[0]

	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", err
	}

	switch typeOfRepeat {
	case "d":
		if len(arr) < 2 {
			return "", errors.New("no interval for d")
		}

		interval, err := strconv.Atoi(arr[1])
		if err != nil {
			return "", errors.New("interval conversion error")
		}

		if interval > 400 || interval == 1 {
			return "", errors.New("interval number isn't allowed (1 <= interval < 400)")
		}

		for {
			date = date.AddDate(0, 0, interval)
			if date.After(now) {
				break
			}
		}

	case "y":
		if len(arr) != 1 {
			return "", errors.New("no additional parameters allowed")
		}
		for {
			date = date.AddDate(1, 0, 0)
			if date.After(now) {
				break
			}
		}
	default:
		return "", errors.New("repetition type not valid")
	}

	return date.Format("20060102"), nil

}

func nextDayHandler(w http.ResponseWriter, req *http.Request) {

	now_req := req.FormValue("now")
	date_req := req.FormValue("date")
	repeat_req := req.FormValue("repeat")

	now, err := time.Parse("20060102", now_req)
	if err != nil {
		http.Error(w, "error while parsing parameter", http.StatusBadRequest)
		return
	}
	if date_req == "" || repeat_req == "" {
		http.Error(w, "date or repeat parameters missing", http.StatusBadRequest)
		return
	}
	res, err := NextDate(now, date_req, repeat_req)
	if err != nil {
		http.Error(w, "error while calculating next date", http.StatusBadRequest)
		return
	}
	w.Write([]byte(res))

}
