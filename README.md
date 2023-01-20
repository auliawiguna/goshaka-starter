Inspired by [Fiber backend template for Create Go App CLI](https://github.com/create-go-app/fiber-go-template) ^^

## Setup
- To run development mode, execute `make watch`
- [Linux] Install `swag` by execute `go install github.com/swaggo/swag/cmd/swag@latest`
- Generate swagger docs by execute `make swagger`

## Features
- Login
- Logout
- Register
- Verify Registration
- Reset Password
- My Profile
- Change Profile
- Role Based Access Control
- Example of CRUD function
- Verification on email change

## Capabilities
- Auto migration
- Auto seeder

## Security Rules
- In this boilerplate, we use stateless JWT
- To change your email address, you need to verify the changes via email