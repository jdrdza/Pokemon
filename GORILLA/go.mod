module GORILLA

go 1.18

replace Pokemon/constants => ../constants

replace Pokemon/csv => ../csv

replace Pokemon/rowsOperations => ../rowsOperations

replace Pokemon/structs => ../structs

require (
	Pokemon/constants v0.0.0-00010101000000-000000000000
	Pokemon/csv v0.0.0-00010101000000-000000000000
	Pokemon/rowsOperations v0.0.0-00010101000000-000000000000
	Pokemon/structs v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)
