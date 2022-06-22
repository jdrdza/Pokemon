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
    <li> In the terminal, go to the cloned repository and execute the following command</li>
  
```
go run main.go
```
</li>
<li>This will have deployed the server using the GIN framework</li>
</ol>
<br>
<h2>3- Enabled endpoints</h2>
<ol>
<li>GET pokemon
    <ol>
      <li>This endpoint returns a list of all pokemon from a CSV file. The way to call this endpoint is the following
        
```
http://localhost:8080/pokemon
```

The response will be displayed as a JSON such as the following
      
        
```
{
    "count": 3,
    "pokemon": [
        {
            "id": 1,
            "name": "bulbasaur",
            "region": "kanto"
        },
        {
            "id": 2,
            "name": "ivysaur",
            "region": "kanto"
        },
        {
            "id": 252,
            "name": "treecko",
            "region": "hoenn"
        }
    ]
}
```
        
</li>
  </ol>

<li>GET pokemon by region
    <ol>  
      <li> This endpoint returns only the pokemon by an specific region. The endpoint is the following
  
  
```
http://localhost:8080/pokemon/region/:region
```

where <<:region>> is the parameter to search. These are all the possible regions

```
kanto
johto
hoenn
sinnoh
teselia
kalos
alola
galar
```
 
The way to call this endpoint is the following
        
```
http://localhost:8080/pokemon/region/kanto
```     

and the response is 
        
```
{
    "count": 2,
    "pokemon": [
        {
            "id": 1,
            "name": "bulbasaur",
            "region": "kanto"
        },
        {
            "id": 2,
            "name": "ivysaur",
            "region": "kanto"
        }
    ]
}
``` 
</li>
  </ol>
  

<li>GET pokemon by ID
    <ol>  
      <li> This endpoint returns the pokemon by its ID. The endpoint is the following
  
  
```
http://localhost:8080/pokemon/id/:id
```

where <<:id>> is the parameter to search. The way to call this endpoint is the following
        
```
http://localhost:8080/pokemon/id/252
```     

and the response is 
        
```
{
    "count": 1,
    "pokemon": [
        {
            "id": 252,
            "name": "treecko",
            "region": "hoenn"
        }
    ]
}
``` 
</li>
  </ol>  
  
  
<li>GET pokemon by name
    <ol>  
      <li> This endpoint returns the pokemon by its name. The endpoint is the following
  
  
```
http://localhost:8080/pokemon/name/:name
```

where <<:name>> is the parameter to search. The way to call this endpoint is the following
        
```
http://localhost:8080/pokemon/name/ivysaur
```     

and the response is 
        
```
{
    "count": 1,
    "pokemon": [
        {
            "id": 2,
            "name": "ivysaur",
            "region": "kanto"
        }
    ]
}
``` 
</li>
  </ol>
  
  <li>GET pokemon by type
    <ol>  
      <li> This endpoint returns the odd and even pokemon. The endpoint is the following
  
  
```
http://localhost:8080/pokemon/types/
```

This endpoint has three query parameters
        
```
type: it can only have the word "odd" or "even"
items: number the items to display in the response
items_per_worker: number of items it will have the workers inside the program
```

The way to call this endpoint is the following

```
http://localhost:8080/pokemon/types/?type=even&items=2&items_per_worker=1
```

and the response is 
        
```
{
    "count": 2,
    "pokemon": [
        {
            "id": 2,
            "name": "ivysaur",
            "region": "kanto"
        },
        {
            "id": 252,
            "name": "treecko",
            "region": "hoenn"
        }
    ]
}
``` 
</li>
  </ol>
  
<li>POST new pokemon
    <ol>  
      <li> This endpoint creates a new pokemon if not exists. The endpoint is the following
  
  
```
http://localhost:8080/pokemon/new/
```

This endpoint has the following body
        
```
{
    "id":152,
    "name":"chikorita",
    "region":"johto"
}
```

The way to call this endpoint is the following

```
URL: http://localhost:8080/pokemon/new
        
BODY:
{
    "id":152,
    "name":"chikorita",
    "region":"johto"
}        
```

and the response is 
        
```
{
    "count": 1,
    "pokemon": [
        {
            "id": 152,
            "name": "chikorita",
            "region": "johto"
        } 
    ]
}
``` 
</li>
  </ol>
 
  
 
<li>POST external pokemon
    <ol>  
      <li> This endpoint gets the pokemon from an external API depending the region and it saves them in a CSV file. The endpoint is the following
  
  
```
http://localhost:8080/pokemon/region/:region
```

where <<:region>> is the parameter to search. The regions that can be obtained are in the point 2 from this section. The way to call this endpoint is the following

```
URL: http://localhost:8080/pokemon/region/kanto
        
BODY: empty      
```

and the response is 
        
```
{
    "count": 151,
    "pokemon": [
        {
            "id": 1,
            "name": "bulbasaur",
            "region": "kanto"
        },
        {
            "id": 2,
            "name": "ivysaur",
            "region": "kanto"
        },
        {
            "id": 3,
            "name": "venusaur",
            "region": "kanto"
        },
        
        .
        .
        .
        
        {
            "id": 151,
            "name": "mew",
            "region": "kanto"
        }
    ]
}
``` 
</li>
  </ol> 
 </ol>
  