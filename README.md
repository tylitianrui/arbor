# Arbor
Arbor is an embedded profiling database engine.

## usage

```go
package main
db, err := arbor.Open("./profiles")
defer db.Close()

db.IngestFile(ctx, "cpu.pb.gz")
// fake code to show how to query the database
top, _ := db.Top(ctx, query)
fg, _ := db.Flamegraph(ctx, query)
```

