# pwnpass

CLI tool for determining whether or not a password has been involved in a data breach.


### Install
`go get https://github.com/masonj88/pwnpass` <br>
If Go environment is default, navigate to ~/go/src/github.com/masonj88/pwnpass and: <br>
`go install pwnpass.go`

### Usage
`pwnpass` will ask you for a password and return a string with information on whether or not it has been pwned
according to https://haveibeenpwned.com.

Additionally, you can batch process a set of passwords using the `-batch="path/to/file.txt"` flag.  Returning from the API is slow,
so huge files will take a long time.

Batch processing means the passwords will be in plain text when you pass them in, and the CLI will return information in plain text. 
DO NOT USE FOR PASSWORDS YOU ARE CURRENTLY USING.  Please use `pwnpass` with no arguments and type in your password that way.
