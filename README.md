# proxy

A simple proxy program.

# build

go get github.com/GScalpel/proxy

# Usage

go run index.go -help

  -host string
    	need input host (default "127.0.0.1")
      
  -port string
    	need input port (default "1080")
      
  -sHost string
    	index host (default "8.8.8.8")
      
  -sPort string
    	index port (default "8888")
      
  -select server or local
    	select  to open server or local service(server or local) (default "local")

# Example

server : 
  go run index.go -select server -host x.x.x.x -port xxxx
 
local :
  go run index.go -select local -host 127.0.0.1 -port xxxx -sHost x.x.x.x -sPort xxxx
