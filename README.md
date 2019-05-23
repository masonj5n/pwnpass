# pwnpass
[![Go Report Card](https://goreportcard.com/badge/github.com/masonj88/pwnpass)](https://goreportcard.com/report/github.com/masonj88/pwnpass)

CLI tool for determining whether or not a password has been involved in a data breach.


### Install
`go get github.com/masonj88/pwnpass` <br>
If Go environment is default, navigate to ~/go/src/github.com/masonj88/pwnpass and: <br>
`go install pwnpass.go`

### Usage
`pwnpass` with no arguments will ask you for a password and return a string with information on whether or not it has been pwned
according to https://haveibeenpwned.com.  This is the preferred usage for passwords in use as the password is never displayed or transmitted in plain text.

`pwnpass` with the `-p=password` flag can be used to specify a single password to be checked.

Additionally, you can batch process a set of passwords using the `-batch="path/to/file.txt"` flag.  The file of passwords should be newline delimited. Returning from the API is slow, so files with many passwords will take a long time.

Batch processing passwords, or specifying a password with the `-p` flag will show the passwords in plain text, so 
DO NOT USE FOR PASSWORDS YOU ARE CURRENTLY USING.
Please use `pwnpass` with no arguments and type in your password that way.
