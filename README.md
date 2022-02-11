
# Dev Workflow

    # Pre reqs
    Clone the repository from https://github.com/ambarish-guntupalli/pet-store
    Install Docker

    # build the image (From the project root)
    docker build -t petstore-app
    <img width="1677" alt="build_image" src="https://user-images.githubusercontent.com/23161171/153547619-43303c65-1e04-4eea-badc-79d682c9863f.png">
    <img width="700" alt="image_list" src="https://user-images.githubusercontent.com/23161171/153547631-9e0cb68c-578b-443e-9fa5-14540a8c91e8.png">


    # Run the app
    docker run -d -p 8080:8080 petstore-app
    <img width="656" alt="docker_run" src="https://user-images.githubusercontent.com/23161171/153547660-c91459bb-e7a6-4754-881c-b8b4e4ae7856.png">
    <img width="1377" alt="docker_ps" src="https://user-images.githubusercontent.com/23161171/153547677-e41ca502-c9d0-4a66-b03c-65527775f166.png">


    # curl commands to test the api
    GET: 
    curl -H "Content-Type: application/json" -X GET 'http://localhost:8080/pets/1'
    <img width="828" alt="get_pet" src="https://user-images.githubusercontent.com/23161171/153547551-6b56651d-73e9-4583-b166-047cf94a3a06.png">


    LIST: 
    curl -H "Content-Type: application/json" -X GET 'http://localhost:8080/pets/'
    <img width="817" alt="list_pets" src="https://user-images.githubusercontent.com/23161171/153547571-6878aa54-1306-4009-b715-fb130aa9243b.png">


    CREATE: 
    curl -H "Content-Type: application/json" -X POST 'http://localhost:8080/pets/' \
    -d '{"id":5, "type":"ox", "price":100}'
    <img width="939" alt="create_pet" src="https://user-images.githubusercontent.com/23161171/153547593-312dfc7c-321a-4303-aa93-ee439bac606a.png">
