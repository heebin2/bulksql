# bulksql

Effective control of sql strings for bulk sql operations.

bulksql automatically increments the argument.

## feature

Clear the front of the DML. (for sqlc)

postgresql, oracle, mysql(sqlite)

```
    (?, ?)   -> (?,?),(?,?),(?,?)
    (:1, :2) -> (:1,:2),(:3,:4),(:5,:6)
    ($1, $2) -> ($1,$2),($3,$4),($5,$6)
```


## example

```go
    bc, err := bulksql.NewBatch("INSERT INTO table_name(att1, att2) VALUES (?, ?) ON ~")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		if err := bc.Push(1, 2); err != nil {
			panic(err)
		}
	}

	sql, err := bc.Query()
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(sql, bc.Datas()...); err != nil {
		panic(err)
	}
```