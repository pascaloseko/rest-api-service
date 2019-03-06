# rest-api-service
Simple rest api implementation in go using database/sql and postgres

**Show Person**
----
  Returns json data about a single person.

* **URL**

  /persons/:uuid

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `uuid=[string]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ id : 1, uuid: 971cd90c-9902-4b21-b9d4-b729f1adf7d8, name : "Michael Bloom", age: 23, created_at: 2019-03-06T12:12:48.089847Z }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "Person with the uuid does not exist" }
    
**Show Persons**
----
  Returns json data of all persons in the database.

* **URL**

  /persons

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
  None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ id : 1, uuid: 971cd90c-9902-4b21-b9d4-b729f1adf7d8, name : "Michael Bloom", age: 23, created_at: 2019-03-06T12:12:48.089847Z }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "Unable to encode response" }`
    
**Add Person**
----
  Returns json data about a single person.

* **URL**

  /persons

* **Method:**

  `POST`
  
*  **URL Params**

   **Required:**
 
  None

* **Data Params**

  ```json
      {
        "name": "Joe Cole",
        "age": 28
      }
  ```

* **Success Response:**

  * **Code:** 201 <br />
    **Content:** `{ id : 1, uuid: 971cd90c-9902-4b21-b9d4-b729f1adf7d8, name : "Michael Bloom", age: 23, created_at: 2019-03-06T12:12:48.089847Z }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "cannot add person" }`
    
**Update Person**
----
  Apply update on a single person.

* **URL**

  /persons/:uuid

* **Method:**

  `PUT`
  
*  **URL Params**

   **Required:**
 
   `uuid=[string]`

* **Data Params**

  ```json
      {
        "name": "Kendric Lamar",
        "age": 28
      }
  ```

* **Success Response:**

  * **Code:** 202 <br />
    **Content:** `{ id : 1, uuid: 971cd90c-9902-4b21-b9d4-b729f1adf7d8, name : "Michael Bloom", age: 23, created_at: 2019-03-06T12:12:48.089847Z }`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "Invalid json in request" }`
    
    
**DELETE Person**
----
  Apply delete on a single person.

* **URL**

  /persons/:uuid

* **Method:**

  `DELETE`
  
*  **URL Params**

   **Required:**
 
   `uuid=[string]`

* **Data Params**

  None

* **Success Response:**

  * **Code:** 204 <br />
    **Content:** 
    None
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "Cannot delete person" }`
