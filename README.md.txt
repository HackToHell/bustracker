#Open Source GPS Bus tracker

##Install
Download and install Go http://golang.org/dl/
Now download an IDE of your choice like Atom from http://atom.io or LiteIDE

We now need to get the pq package and also the protobuf package, before that we need to set up the environment for go called "Workspace"

Make a folder where the packages you pull will be placed and add it to the path as GOPATH(different instructions for different platforms, use google)

Now make sure you have git and mercurial executables in your path before getting the above mentioned packages as go simply does a clone of the repository

    go get github.com/lib/pq

For protobuf 
    go get code.google.com/p/goprotobuf/{proto,protoc-gen-go}

After installing protobuf from http://code.google.com/p/protobuf/

Finally run

    go run server.go

This should start the protobuf server, the HTTP server is server2.go(Still incomplete). Currently uses Postgres, will add schema later.  
