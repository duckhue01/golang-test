##  Todo list API using Golang with protobuf, and grpc-gateway



### Additional Requirements:
  1. main idea: i use enum type in Protocol Buffers to present 3 states of todo.

  2. main idea: i create extra table to store Tags of Todo, this will prevent database from store redundant data.

  3. main idea: i create sql query for 4 scenarios:<br>
    - status == nil && tags == nil<br>
    - status != nil && tags == nil<br>
    - status == nil && tags != nil<br>
    - status != nil && tags != nil <br>
  *but i feed that this solution is'nt the best, because it too long*<br>

  4. main idea. 
    - i create new field in table in order to store the index of Todo  and take in two input:
      1. Id: id of Todo
      2. Pos: index that we want to move to
    - There are 2 cases:
      1. bot-to-top: we want move Todo from larger index to smaller index
      2. top-to-bot: we want to move Todo from smaller index to larger index
    
    - in case bot-to-top:
      1. we increase all index of Todo in range [Pos, Id) by 1
      2. we change the index of Todo to Pos

    - in case top-to-bottom:
      1. we decrease all index of Todo in range (Id, Pos] by 1
      2. we change the index of Todo to Pos


### Diagram:
![](dia.png)

### Setup with Docker:
  **Prerequisites**
  1. Docker & Docker Compose
  2. Git

  **Setup Instructions**
  1. Clone the project to local machine and go to the folder<br>
    ```
    git clone https://github.com/duckhue01/golang_test.git
    ```
    <br>
    ```
      cd ./golang_test
    ```
  2. Use docker compose to build images and run containers<br>
    ```
    docker-compose up --build
    ```
### REST API Documentation:
  **Create Todo**
  ----
  Create a new todo item.
  * **URL** `/v2/todos`
  * **Method:** `POST`
    
  * **Success Response:**
    * **Code:** 200 OK <br />
      ```json
      {
        "api": "v2"
      }
      ```
  
  * **Sample Call:**

    ```json
    POST  http://127.0.0.1:8080/v2/todos HTTP/1.1
    Content-Type: application/json

    {
      "todo":{
          "id":1,
          "title":"dasdas",
          "description":"asdasdasd",
          "createAt":"2020-01-15T01:30:15.01Z",
          "updateAt":"2020-01-15T01:30:15.01Z",
          "status":0,
          "tags":["sleep", "relax"]
      }
    }
    ```

  **Get All**
  ----
  Get all todos are stored in database
  * **URL** ` /v2/todos`
  * **Method:** `GET`
    
  * **Success Response:**

    * **Code:** 200 OK <br />  
      ```json
      {
        "api": "v2",
        "todo": [
          {
            "id": 2,
            "title": "dasdas",
            "description": "asdasdasd",
            "createAt": "2020-01-15T01:30:15Z",
            "updateAt": "2020-01-15T01:30:15Z",
            "tags": [
              "relax",
              "sleep"
            ],
            "status": "DONE",
            "order": 2
          },
          {
            "id": 1,
            "title": "dasdas",
            "description": "asdasdasd",
            "createAt": "2020-01-15T01:30:15Z",
            "updateAt": "2020-01-15T01:30:15Z",
            "tags": [
              "relax",
              "sleep"
            ],
            "status": "DONE",
            "order": 1
          }
        ]
      }
      ```
  * **URL parameters:** 
    pag,tags,status




  * **Sample Call:**

    ```json
    GET http://127.0.0.1:8080/v2/todos?pag=100&tags=Sleep&tags=relax&status=2  HTTP/1.1
    Content-Type: application/json
    ```


    
  **Get One**
  ----
  get one todos with id
  * **URL** `/v2/todos/:id` 
 
  
  * **Method:** `GET`
    
  * **Success Response:**

    * **Code:** 200 OK <br />  
      ```json
      {
        "api": "v2",
        "todo": {
          "id": 2,
          "title": "dasdas",
          "description": "asdasdasd",
          "createAt": "2020-01-15T01:30:15Z",
          "updateAt": "2020-01-15T01:30:15Z",
          "tags": [
            "relax",
            "sleep"
          ],
          "status": "DONE",
          "order": 3
        }
      }
      ```
      * **Error Response:**

    * **Code:** 404 NOT FOUND <br />
      ```json
      {
        "code": 5,
        "message": "Todo with ID='12' is not found",
        "details": []
      }
      ``` 


  * **Sample Call:**

    ```json
    GET http://127.0.0.1:8080/v2/todos/12 HTTP/1.1
    ```


  **Update Todo**
  ----
  update existing todo
  * **URL** `/v2/todos/`
  
  
  * **Method:** `PUT`
    
  * **Success Response:**

    * **Code:** 200 OK <br />  
      ```json
      {      
        "api": "v2"
      }
      ```
  * **Error Response:**

    * **Code:** 404 NOT FOUND <br />
      ```json
      {
        "code": 5,
        "message": "Todo with ID='12' is not found",
        "details": []
      }
      ``` 
  
  * **Sample Call:**
    ```json
    PUT  http://127.0.0.1:8080/v2/todos HTTP/1.1
    Content-Type: application/json

    {
      "api": "v2",
      "todo": {
        "id": 2,
        "title": "duckhue01",
        "description": "duckhue01",
        "createAt": "2020-01-15T01:30:15Z",
        "updateAt": "2020-01-15T01:30:15Z",
        "tags": [
          "relax",
          "sleep"
        ],
        "status": "TODO",
        "order": 3
      }
    }
    ```


  **Delete Todo**
  ----
  delete existing todo
  * **URL** `/v2/todos/:id`
  
  
  * **Method:** `DELETE`
    
  * **Success Response:**

    * **Code:** 200 OK <br />  
      ```json
      {      
        "api": "v2"
      }
      ```
  * **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    ```json
    {
      "code": 5,
      "message": "Todo with ID='12' is not found",
      "details": []
    }
    ``` 
  * **Sample Call:**
    ```json
    DELETE http://127.0.0.1:8080/v2/todos/1 HTTP/1.1
    ```



  **Reorder Todo**
  ----
  reorder existing todo
  * **URL** `/v2/todos/reorder/:id`
  
  
  * **Method:** `PUT`
    
  * **Success Response:**

    * **Code:** 200 OK <br />  
      ```json
      {      
        "api": "v2"
      }
      ```
  * **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    ```json
    {
      "code": 5,
      "message": "Todo with ID='12' is not found",
      "details": []
    }
    ``` 
  * **Sample Call:**
    ```json
    PUT  http://127.0.0.1:8080/v2/todos/reorder/1 HTTP/1.1
    Content-Type: application/json

    {
      "api":"v2",
      "pos":10
    }
    ```

    ```json
    PUT  http://127.0.0.1:8080/v2/todos/reorder/10 HTTP/1.1
    Content-Type: application/json

    {
      "api":"v2",
      "pos":1
    }
    ```