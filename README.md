# redis web app for testing

To build with Docker:

```
docker build -t rediweb:latest .
```

To run:

```
docker run -P rediweb:latest
```

To build and run with Go:

```
go get -v -u github.com/schmichael/rediweb
rediweb
```
