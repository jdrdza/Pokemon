# Pokemon

<br>
<h2>1.- Instructions for clonning the repository</h2>
<ol>
    <li>Open a terminal and run the following command:

```
git clone https://github.com/jdrdza/Pokemon.git
```
</li>
</ol>

<br>
<h2>2.- Instructions for Running</h2>
<ol>
    <li> In the terminal, go to the cloned repository and enter the main folder</li>

```
cd/main
```
</li>
<li> In the main folder execute the following command

```
go run main.go
```
</li>
<li>This will have deployed the server using the GIN framework. However, if you want to change this framework to ECHO or Gorilla, in the main function of main.go go, you can change the function

```
server("gin")
```

to

```
server("echo")
```
or

```
server("gorilla")
```
</li>
</ol>
<br>
<h2>3- Enabled endpoints</h2>
<ol>
<li>These are the endpoints that you can use to get the pokemon information

```
http://localhost:8080/allPokemon
http://localhost:8080/pokemonByRegion/:region
http://localhost:8080/pokemonById/:id
http://localhost:8080/pokemonByName/:name
```
</li>
</ol>
