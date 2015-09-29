Journalist
==========

To run
------

    cd Journalist/
    export GOPATH=`pwd`
    cd src/
    go run main.go <file>

Where <file> is an IRC log.

Formats
-------

Currently it only understands _very simple_ IRC logs.  Only messages.

Languages
---------

Currently it only exports to Danish and is hardcoded as such.

Roadmap
-------

Should eventually support more IRC log messages, more languages and hopefully
be a web server in addition to being a command line program.
