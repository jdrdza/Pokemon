package router

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jdrdza/pokemon/control"
	"github.com/jdrdza/pokemon/file"
	"github.com/jdrdza/pokemon/method"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Server   string `yaml:"server"`
	EpBase   string `yaml:"ep-base"`
	Gregion  string `yaml:"g-region"`
	Gid      string `yaml:"g-id"`
	Gname    string `yaml:"g-name"`
	Gtype    string `yaml:"g-type"`
	Pnew     string `yaml:"p-new"`
	Pregion  string `yaml:"p-region"`
	Prestart string `yaml:"p-restart"`
}

type router struct {
	control control.Controller
}

type Router interface {
	Routers()
}

func newRouter(control control.Controller) *router {
	return &router{control: control}
}

func Initialise() (route *router, err error) {
	rows, err := file.NewFile()
	if err != nil {
		return nil, err
	}
	meth := method.NewMethod(rows)
	cont := control.NewController(meth)
	route = newRouter(cont)
	return route, err
}

func config() (*conf, error) {
	data, err := ioutil.ReadFile("data\\config.yaml")
	if err != nil {
		return nil, err
	}
	y := &conf{}

	err = yaml.Unmarshal([]byte(data), &y)
	if err != nil {
		return nil, err
	}

	return y, nil
}

func (route *router) Routers() {

	config, err := config()
	if err != nil {
		log.Println(err.Error())
		return
	}

	router := gin.Default()
	defer router.Run(config.Server)

	post := router.Group(config.EpBase)
	{
		post.POST(config.Pnew, func(ctx *gin.Context) {
			route.control.NewPokemon(ctx)

			route, err = Initialise()
			if err != nil {
				log.Println("There was an error: " + err.Error())
				return
			}
		})

		post.POST(config.Pregion, func(ctx *gin.Context) {
			route.control.PokeAPI(ctx)

			route, err = Initialise()
			if err != nil {
				log.Println("There was an error: " + err.Error())
				return
			}
		})

		post.POST(config.Prestart, func(ctx *gin.Context) {
			route, err = Initialise()
			if err != nil {
				log.Println("There was an error: " + err.Error())
				return
			}

			ctx.IndentedJSON(http.StatusOK, "The server was restarted")
		})

	}

	get := router.Group(config.EpBase)
	{
		get.GET("/", func(ctx *gin.Context) {
			route.control.AllPokemon(ctx)
		})

		get.GET(config.Gregion, func(ctx *gin.Context) {
			route.control.PokemonByRegion(ctx)
		})

		get.GET(config.Gid, func(ctx *gin.Context) {
			route.control.PokemonById(ctx)
		})

		get.GET(config.Gname, func(ctx *gin.Context) {
			route.control.PokemonByName(ctx)
		})

		get.GET(config.Gtype, func(ctx *gin.Context) {
			route.control.Items(ctx)
		})
	}

}
