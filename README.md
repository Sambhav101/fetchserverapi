# Fetch Server API
It is a simple API built on Go that processes receipts and generates points for each receipt.

---


## API Specification
#### Endpoint: Process Receipts
- Path: ```/receipts/process```
- Method: ```POST```
- Payload: Receipt JSON
- Response: JSON containing an id for the receipt.
- Description: Takes in a JSON receipt (see example in the example directory) and returns a JSON object with a unique ID

##### Example Response:
```
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

#### Endpoint: Get Points

- Path: ```/receipts/{id}/points```
- Method: ```GET```
- Response: A JSON object containing the number of points awarded.
- Description: A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

##### Example Response:
```
{ "points": 32 }
```

---

## Dependencies
1. ```github.com/google/uuid``` - Generates unique ID at all times (even for the same receipt json)
2. ```github.com/gorilla/mux``` - Framework to create HTTP routes.

---

## How to Run

1. Clone the repo and ```cd``` into the directory.

2. If you have Go installed, skip step 3, else go to step 3
   ```
   go run main.go
   ```

3. If you have docker installed
  - Build the docker image
    ```
    docker build -t your-app-name .
    ```
  - Run the docker image
    ```
    docker run -p 8080:8080 your-app-name
    ```

4. Open a new terminal
  - For ```POST``` request
    ```
    curl -X POST -d '{your receipt json}' http://localhost:8080/receipts/process
    ```
    or if you are inside the repo
    ```
    curl -X POST -d '@examples/simple-receipt.json' http://localhost:8080/receipts/process
    ```
 - For ```GET``` request
   ```
   curl http://localhost:8080/receipts/{generated id from post request}/points
   ```
   or go to the same link in your web browser



