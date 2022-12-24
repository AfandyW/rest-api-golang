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
	    "code": "201",
	    "status": "Success",
	    "message" : "Note Baru Berhasil Ditambahkan"
    }
    ```