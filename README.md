## Todo API in GO

This public repo demonstrate a todo app API development using Go lang.

## Library / Packages Used

- Chi (router) - github.com/go-chi/chi/v5
- Jwt (From Chi) - github.com/go-chi/jwtauth
- Gorm ( ORM ) - gorm.io/gorm
- Render ( response helper ) - github.com/go-chi/render
- Bcrypt ( Password Encoder / Decoder ) -  golang.org/x/crypto
- Postgres DB Driver - gorm.io/driver/postgres


### Features include.

Todo:

- [x] Create Todo 
- [x] List All Todos 
- [x] Edit Todo 
- [x] Delete Todo 
- [x] Toggle Todo Is Completed / Not Completed
- [x] Attach User ID to Todos.
- [x] Prevent Unauthorized users from creating Todo
  

Users:

- [x] Create New User 
- [x] Login User
- [x] JWT Token Generation

