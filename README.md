# Setup

First initialize the db to create the `.db` folder to persist the DB information in the local environment (only nedd in the first run).

```
docker-compose up postgres -d
```

Now you can run the application to execute the migration and start the server.

```
docker-compose up
```

