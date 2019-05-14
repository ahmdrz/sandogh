### Sandogh
> Simple Image storage for Fandogh PaaS with cache support for your browsers.

##### Available HTTP methods

The format of URi is like `/files/:directory/:name`. So if your domain is for example `fandogh.cloud` your final URL is `fandogh.cloud/files/:directory/:name`. In any cases, `Authentication` header is your secret key.

1. GET an image from your storage.
2. POST an image to your storage (Authentication header needed).
3. DELETE an image from your storage (Authentication header needed).

##### Local installation using Docker

```
$ docker pull ahmdrz/sandogh:latest
$ docker run -p 8080:8080 \
             -e SERVICE_SECRET_KEY='<storage secret key>' \
             -e SERVICE_BASE_DIRECTORY='/var/storage' \
             -v storage:/var/storage ahmdrz/sandogh:latest 
```

##### Using Fandogh

```
$ fandogh login --username '<your username>' --password '<your password>'
$ fandogh namespace active --name '<your namespace>'
$ fandogh volume add --name 'images' --capacity 10
$ fandogh service apply -f service.yml --secret_key='<storage secret key>'
```
