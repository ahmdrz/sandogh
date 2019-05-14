### Sandogh
> Simple Image storage for Fandogh PaaS 

##### Local installation using Docker

```
$ docker pull ahmdrz/sandogh:latest
$ docker run -p 8080:8080 \
             -e SERVICE_SECRET_KEY='18561b13c56b' \
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