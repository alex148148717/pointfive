# raw events


## Technologies Used

- Go (Golang)
- PostgreSQL



```bash
./app 
```


## Importing job


```bash

curl -X POST http://localhost:8081/import_job/v1/job \
  -F "file=@examples/data.txt"
```
