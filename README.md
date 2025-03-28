# go-basics

#### https://github.com/mschwarzmueller/go-complete-guide-resources
#### https://pkg.go.dev/std
 
### commands
```
go run app.go
```

### turn project into module
```
go mod init example.com/first-app
```
```
go build
```

### Go comes with a couple of built-in basic types:

1. int => A number WITHOUT decimal places (e.g., -5, 10, 12 etc)
2. float64 => A number WITH decimal places (e.g., -5.2, 10.123, 12.9 etc)
3. string => A text value (created via double quotes or backticks: "Hello World", `Hi everyone')
4. bool => true or false

## To run a go function in dev mode
1. ```go run investment_calculator.go``` 
2. ```go run .``` -> when go.mod is present

#### Create a new module
1. Create a new folder named `first-app`
2. run command `go mod init example.com/first-app`