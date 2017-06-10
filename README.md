# Go Example for the error interface / how to handle and identify error types

This little example has two error types which support the error interface. 100 errors are generated randomly, the error type will pe checked and printed out.

To run this example in linux use following:

```{r, engine='bash', count_lines}
go get github.com/devilsray/golang-errorinterface-example
cd $GOPATH/src/github.com/devilsray/golang-errorinterface-example/
go run *.go
```

The output will look like

0: this is an error from type TypeAError

1: and this is an error from type TypeBError

2: this is an error from type TypeAError

3: this is an error from type TypeAError

.....
