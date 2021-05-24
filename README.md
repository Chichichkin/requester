# Go Requester
This is simple program that allow to send GET requests to different sites and 
return string? that consist of url and md5 hash of body response
## How to use
For example, you can run this program by calling:

    go run main.go -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com

This program by default runs with 10 goroutines, but that number can be changed this way:

    -p 3
    //or
    -parallel 3

You can also use file,from which program will read sites addresses:

    -f sites.txt
    //or 
    -file sites.txt
***Currently supports files where only one site per line***

So final command will look like this:

    go run main.go -parallel 3 -f sites.txt
## Test
To run test to check if everything working correct, please run:

    go test -v
In project directory.