package usecases

import (
	"errors"
	"net/http"
	"strconv"
)

type UseCase struct {
}

// New community UseCase
func New() *UseCase {
	return &UseCase{}
}

func (u *UseCase) WorkerFlags(r *http.Request) (string, int, int, error) {
	type_arr, ok := r.URL.Query()["type"]
	var err error
	var type_str = ""
	var items = -1
	var items_per_worker = -1
	if !ok || len(type_arr) < 1 {
		return "", 0, 0, errors.New("Url Param 'type' is not given")
	}
	if len(type_arr) > 0 {
		type_str = type_arr[0]
		if type_str != "odd" && type_str != "even" {
			return "", 0, 0, errors.New("Only support 'odd' or 'even' types")
		}
	}

	items_arr, ok := r.URL.Query()["items"]
	if !ok || len(items_arr) < 1 {
		return "", 0, 0, errors.New("Url Param 'items' is not given")
	}
	if len(type_arr) > 0 {
		items_str := items_arr[0]
		items, err = strconv.Atoi(items_str)
		if err != nil {
			return "", 0, 0, errors.New("Items must be an int")
		}
	}

	items_per_worker_arr, ok := r.URL.Query()["items_per_worker"]
	if !ok || len(items_per_worker_arr) < 1 {
		return "", 0, 0, errors.New("Url Param 'items_per_worker' is not given")
	}
	err = nil
	if len(items_per_worker_arr) > 0 {
		items_per_worker_str := items_per_worker_arr[0]
		items_per_worker, err = strconv.Atoi(items_per_worker_str)
		if err != nil {
			return "", 0, 0, errors.New("items_per_worker must be an int")
		}
	}
	return type_str, items, items_per_worker, nil
}
