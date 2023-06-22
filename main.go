package main

import (
	sq "github.com/SG143s/uaswrpl/sqlop"
	web "github.com/SG143s/uaswrpl/webht"
)

func main() {
	sq.Sqinit()
	web.Webstart()
}
