** Install and Config **

*** Configing the workspace ***

Copy the files bloomfilter.go and bloomfilter_test.go in a directory

Next execute the following commands (in the directory)

$ go mod init katas/bloomfilter
$ go work use .
$ go mod tidy

The solution needs the package **github.com/datadog/mmh3**

$ go test -v
