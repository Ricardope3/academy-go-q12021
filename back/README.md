- To get all data in the csv

`http://localhost:8000/pokemons`

- To get a specific pokemone using an ID

`http://localhost:8000/pokemons?id=3`

- To get All todos in the JSON placeholder api

`http://localhost:8000/todos`

```data.csv``` will be saved locally in the root folder of the project
containing all of the todos fetched from JSON placeholder. 

- To make the workers search for pokemons concurrently

`http://localhost:8000/workers?type=even&items=10&items_per_worker=100`

### Query Params
All paremeters are not optional!!!
- type: `even` || `odd` 
- items: `int`
- items_per_worker: `int`
=======
```http://localhost:8000/pokemons?id=3```
