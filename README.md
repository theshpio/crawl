Crawl fetches, sorts and outputs a list of url from the given one. It makes available content easy to find.
>Install binary. [more details](https://go.dev/doc/code).
> Hint if you have a trouble:
```
# navigate to directory and run:
go install .

# follow output

# for linux add to ~/.bashrc:
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

# source it and run again
go install .
```

> usage:
```bash
crawl -url http://go.dev
```
> output:
```
http://go.dev/
http://go.dev/blog
http://go.dev/blog/
http://go.dev/brand
...
```
Additionally you could wrap it with torify.
Follow this clear instructions from the [source](https://justhackerthings.com/post/using-tor-from-the-command-line/)
> usage with torify:
```bash
torify crawl -url http://go.dev
```
