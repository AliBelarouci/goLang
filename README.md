# DATA SCORE CHALLENGE

this my challenge for datascore team.

### you find in this details :

- 1. how to implemente requirements of this project.
- 2. how to run project dev mod or prod mod.
- 3. implementation the test.
- 4. how to configure gcp for reading csv file from it.
- 5. Setup a CI/CD pipeline to build or test.

### Technologies

- Go
- PostgreSQL
- Docker

1. ### implementation of Requirements

- USING gorm ORM and fiber V2 for http server
- implemetation connection for database
- generate connection string by reading .env file
- generate model for the customer
- migrate the model customer for crating table in postegresql server
- create Routers by using "github.com/gofiber/fiber/v2

#### Routes

- **Get /customers** Get all customers saved on database
- **Post /customer** Add new customer from Front end post
- **Post /upload_csv** Add list of customers from csv file
- **Post /upload_and_track_file_csv** Add list of customers from csv file after check if the file already uploaded (track by the name of file)
- **Post /upload_and_track_file_csv_SHA_256** Add list of customers from csv file after check if the file already uploaded by generate a unique hash for a CSV file by reading its contents and compaire this hash by the hash codes that save in database
- **Post /google_cloud_bucket_file_csv** read data form googe storage use bucket and read csv file and save data to database

#### Run the source code

```bash
go run main.go
```

2. ### run project develepment mode or production mode

#### develepment mode

create docker-compose-postegresql.yml file for postgresql database and we use .env file for authentication to database server and run the command bellow

```bash
docker-compose -f "docker-compose-postegresql.yml" up -d
```

after running the postegresql container we run this command for the REST API server

```bash
go run main.go
```

#### production mode

in this method we have build docker image for our project by using **Dockerfile** and use it in docker-compose.yaml
just run thi command

```bash
docker-compose -f "docker-compose.yaml" up -d --build
```

3. ### implementation the test

- implementation test for connection to postegresql database.
- implementation test for create new customer and new file

4. ### Reading file from GCP Storage

#### Create accoute and project

- Go to the GCP website: https://cloud.google.com/
- sign in with an existing Google account or create new one.
- Once you are signed in, you will be prompted to set up a new project. Follow the instructions to complete the setup process.

#### Create Bucket and upload file

1. Navigate to the Cloud Storage section.
2. Click the "Create bucket" button.
3. Enter a unique name for your bucket and select a storage class, location, and default
4. access control settings for our case **alibucket123**.
5. Click the "Create" button to create your bucket.
6. Click on the newly created bucket to open its details.
7. Click the "Upload files" button for our case **test-1.csv**.
8. Select the file you want to upload.
9. Set the access control settings for the file if desired.
10. Click the "Upload" button to upload the file to your bucket.

#### credentials JSON file

To obtain the credentials JSON file, for the access to the bucket and file follow these steps:

1. Navigate to service account
2. Click on the "Create service account" option on the top
3. Fill out the necessary information for the service account
4. After the service created you can manage keys and click on "Add key" and select 'create new key ' choose json and click "create" then you can download the json file and set in our root project ("For our case the file is: credentials.json")
5. ### Setup a CI/CD pipeline to build and test
   in the file **_.gitlab-ci.yml_** sets up a GitLab CI/CD pipeline for a **Go project**. Here are the steps it performs:
6. Define the base image for the pipeline: golang:1.18-alpine3.16
7. Start a service named postgres:14 with an alias of postgres.
8. Define environment variables for the PostgreSQL database configuration:
   - POSTGRES_DB: the database name
   - POSTGRES_USER: the database user name.
   - POSTGRES_PASSWORD: the database user password.
   - DB_SOURCE: the database connection string.
9. Define two stages for the pipeline: build and test.

#### In the build stage:

Run the go build command to build the application.

```bash
go build
```

Store the built artifact in the datascore/challenge path.

### In the test stage:

Run this command fot the test of project

```bash
CGO_ENABLED=0 go test ./...
```
