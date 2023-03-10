# Rest API - Golang

## Model
- Note
    ```json
    {
        "id": 123,
        "title": "Golang",
        "body" : "Golang adalah bahasa yang dibuat oleh google",
    }
    ```
# Endpoint
## Create Note
- URL : `localhost:8000/api/v1/notes`
- Method: `POST`
- request body
    ```json
    {
	    "title": "Golang",
	    "body" : "Golang adalah bahasa yang dibuat oleh google"
    }
    ```
- response 
    ```json
    {
	    "code": 201,
	    "status": "Success",
	    "message" : "Note Baru Berhasil Ditambahkan"
    }
    ```
## List Note
- URL : `localhost:8000/api/v1/notes`
- Method: `GET`
- response 
    ```json
    {
	    "code": 200,
	    "status": "Success",
	    "message" : "Success get list notes",
        "data": [
            {
                "id": 1,
                "title": "golang",
                "body": "bahasa golang dari google"
            },
            {
                "id": 2,
                "title": "js",
                "body": "javascript beda dari java"
            }
        ]
    }
    ```
## Update Note
- URL : `localhost:8000/api/v1/notes?id={id}`
- Method: `PUT`
- request body
    ```json
    {
	    "title": "Golang",
	    "body" : "Golang adalah bahasa yang dibuat oleh google"
    }
    ```
- response 
    ```json
    {
	    "code": 200,
	    "status": "Success",
	    "message" : "Note Berhasil Diubah"
    }
    ```
## Delete Note
- URL : `localhost:8000/api/v1/notes?id={id}`
- Method: `DELETE`
- response 
    ```json
    {
	    "code": 200,
	    "status": "Success",
	    "message" : "Note Berhasil Dihapus"
    }
    ```