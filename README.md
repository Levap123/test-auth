
# test-auth

Тестовое задание на позицию Junior Go разработчик. 2 "наносервиса" :D.




## Run Locally

Clone the project

```bash
  git clone https://github.com/Levap123/test-auth
```

Go to the project directory

```bash
  cd test-auth
```

Start the server

```bash
  docker-compose build
  docker-compose up
```
POST http://localhost:8080/create-user 


GET http://localhost:8080/get-user/{email}


POST http://localhost:5000/generate-salt


