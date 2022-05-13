module main

go 1.18

replace Pokemon/constants => ../constants

require (
	Pokemon/ECHO v0.0.0-00010101000000-000000000000
	Pokemon/GIN v0.0.0-00010101000000-000000000000
	Pokemon/constants v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/gorilla/mux v1.8.0
	github.com/labstack/echo v3.3.10+incompatible
)

require (
	Pokemon/GORILLA v0.0.0-00010101000000-000000000000 // indirect
	Pokemon/csv v0.0.0-00010101000000-000000000000 // indirect
	Pokemon/rowsOperations v0.0.0-00010101000000-000000000000 // indirect
	Pokemon/structs v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220507011949-2cf3adece122 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace Pokemon/csv => ../csv

replace Pokemon/structs => ../structs

replace Pokemon/rowsOperations => ../rowsOperations

replace Pokemon/gin => ../gin

replace Pokemon/GIN => ../GIN

replace Pokemon/ECHO => ../ECHO

replace Pokemon/GORILLA => ../GORILLA
