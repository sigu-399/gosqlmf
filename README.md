#gosqlmf
Fetches results from SQL rows into a map - Go language

##Status
Completed & Tested : v 1.0

## Install
```
go get github.com/sigu-399/gosqlmf
```

##Usage

Update :

Alternatively, you may use the wrapper functions QueryOne & QueryAll, making things easier and your code smaller.

```
  import(
  	"github.com/sigu-399/gosqlmf"
  )

  // Classic database connection and query 

  database, err := sql.Open(`postgres`, `user=foo dbname=bar password=secret`)
  defer database.close()

  ok, row, err := mf.QueryOne(database, "SELECT * FROM person WHERE id = '5'")
}
```

```
  import(
  	"github.com/sigu-399/gosqlmf"
  )

  // Classic database connection and query 

  database, err := sql.Open(`postgres`, `user=foo dbname=bar password=secret`)
  defer database.close()

  rows, err := mf.QueryAll(database, "SELECT * FROM person")
}
```

### Fetching one row
```
  import(
  	"github.com/sigu-399/gosqlmf"
  )

  // Classic database connection and query 

  database, err := sql.Open(`postgres`, `user=foo dbname=bar password=secret`)
  defer database.close()
	
  rowPlayer, err := database.Query(`SELECT name,score FROM player WHERE id = 67`)
	if err != nil {
		panic(err.Error())
	}

  // Fetch into a map
  
	ok, mapPlayer, err := mf.FetchOne(rowPlayer)
	if err != nil {
		panic(err.Error())
	}
  
  // now can you do things like...
  
  if ok {
    fmt.Printf("%d\n", mapPlayer["score"].(int64) )
  }
```

### Fetching all rows
```
  // Classic database connection and query 

  database, err := sql.Open(`postgres`, `user=foo dbname=bar password=secret`)
  defer database.close()
	
  rowPlayers, err := database.Query(`SELECT name,score FROM player WHERE score > 150`)
	if err != nil {
		panic(err.Error())
	}

  // Fetch into a map
  
	mapPlayers, err := mf.FetchAll(rowPlayer)
	if err != nil {
		panic(err.Error())
	}
  
  // now can you do things like...
  
    fmt.Printf("%d\n", mapPlayers[0]["score"].(int64) ) // first
    fmt.Printf("%s\n", mapPlayers[1]["name"].([]uint8) ) // second etc...
```
