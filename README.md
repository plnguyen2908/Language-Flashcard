# Flashcard Project

## Prerequisite

- This project used go version 1.20
- If you haven't downloaded the Docker, please do so
- You also need to download MongoDB

## Run steps

- In your terminal, run "docker-compose up" without the double quote to start the Docker's containers for 3 services
- Then access the frontend interface through port 8081

## Notes

- The API is hosted at localhost:8080 with the following HTTP's methods:
    - GET: 
        - Get all collections' names at "/flash"
        - Get all of the documents of a collection at "/flash/:collectionName"
    - POST: Add a new document to a collection at "/flash/:collectionName"
    - PUT: Update a document of a collection at "/flash/:collectionName"

- The static web is hosted at localhost:8081

- The mongoDB service is hosted at port 27017 of the mongo service of the container

## Copyright

The project is built by Phuc (Leos) Nguyen, any reuse should have Phuc's permission.