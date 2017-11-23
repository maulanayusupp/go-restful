package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlerTodosSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}

	todo := RepoFindTodo(todoId)

	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
