# quick-note 
### fast, secure, persistant notes


## Synopsis

So the service I used for my notes sold out and went too comercial for my taste. After not finding a decent replacement in 15 minutes I decided  I will build my own, because why not ?



Please read chapters Install, Build and Run and Configuration if you want to host your own service. 

If you want to use only the backend API have a look at the chapter "Backend API"


## Install
You need **go** and **nginx** installed. If you have apt run ``bash install.sh`` from the project's directory. 
If you don't you have to install them manually.

## Build
Run ``bash build.sh`` this will run all the tests, verify that you have a correct environment set up and build the binary. If there is an error in this step it usually means that the environment is not set up correctly. (eg GOPATH is not set)

## Run & Configuration

## Backend API
The backend API is RESTful. The Payload of the response contains some useful text that can be shown to the user.

### API Calls
| API Call | HTTP Method | Path              | Request Payload              | Description                                                                                                                                     |
|----------|-------------|-------------------|------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| Save      | POST        | /save/             |Note contents | Sets a note's contents. |
| Get    | GET         | /g/<noteid>  | -                            | Fetches a note's contents.|           

### API Status Codes

| Code                      | Description                                              |
|---------------------------|----------------------------------------------------------|
| 200 OK                    | A request was completed successfully                     |
| 404 Not found             | Method/File doesn’t exist                                |
| 400 Bad Request           | Wrong parameters were passed when making the API Call. Also used for authorization errors.   |
| 500 Internal Server Error | There was an internal error. The request can be retried. |  # quick-note
