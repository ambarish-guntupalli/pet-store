
# Dev Workflow

    # Pre reqs
    Clone the repository from https://github.com/ambarish-guntupalli/pet-store
    Install Docker

    # build the image (From the project root)
    docker build -t petstore-app

    # Run the app
    docker run -d -p 8080:8080 petstore-app

    # curl commands to test the api
    GET: 
    curl -H "Content-Type: application/json" -X GET 'http://localhost:8080/pets/1'

    LIST: 
    curl -H "Content-Type: application/json" -X GET 'http://localhost:8080/pets/'

    CREATE: 
    curl -H "Content-Type: application/json" -X POST 'http://localhost:8080/pets/' \
    -d '{"id":5, "type":"ox", "price":100}'