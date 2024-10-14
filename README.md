# band
Go Library for building micro-services.

## band cmd
- init project
```sh
make project
cd project
band init -p github.com/9d77v/project
```  
  
- add new service
```sh
band service -s project -e task
```

- start service
```sh
go mod tidy
make wire-project
make project-service
```
