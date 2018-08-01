[![](https://images.microbadger.com/badges/image/nbarnum/containerip.svg)](https://microbadger.com/images/nbarnum/containerip "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/nbarnum/containerip.svg)](https://microbadger.com/images/nbarnum/containerip "Get your own version badge on microbadger.com")

# containerip

Go HTTP server returns all of the container's IPv4 addresses.

## Docker

### Use existing image

Run the container:

```
$ docker run --rm -it -p 8080:80 nbarnum/containerip
```

Get the IPs:

```
$ curl http://localhost:8080
172.17.0.2
```

### Build Docker image

Build the image:

```
$ docker build -t $USER/containerip .
```

Clean up intermediate containers left behind from multi-stage Dockerfile:

```
$ docker system prune
```

Run the container:

```
$ docker run --rm -it -p 8080:80 $USER/containerip
```
