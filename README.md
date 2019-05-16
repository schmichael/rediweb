# redis web app for testing

To build with Docker:

```
docker build -t schmichael/rediweb:0.1 .
```

To run:

```
docker run -P schmichael/rediweb:0.1
```

To build and run with Go:

```
go get -v -u github.com/schmichael/rediweb
rediweb
```
