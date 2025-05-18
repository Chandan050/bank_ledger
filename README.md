### run (for local testing)


```
go run cmd/main.go
```
### create new image 

```
docker build -t bank_ledger .
```
### run container's image for postgres, mongodb , kafka , zookeeper and current build

```
docker compose up
```

### ports 

- 5432 for postgres
- 27017 for mongo
- 9092 for kafka
- 8080 for build
</br>

### APi collection contains 
- postman collection of all endpoints (Bank_ledger.postman_collection.json)
- for each endpoint, there is a request body 
- *.rest file contains the request body for each endpoint to test in vscode itself
</br>
    - *Get account                                 /account/:id
    - Create the account                          /account
	- Creates the transaction                     /transaction
	- get's single transaction                    /transactions/:transaction_id
	- return the all transaction of customer      /transactions_list/:account_id