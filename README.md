![image](https://user-images.githubusercontent.com/26473549/214473471-19e5a263-cf21-440a-9beb-d2cc16eed9fc.png)

# A production ready Golang Boilerplate API

Inspired by [Fiber backend template for Create Go App CLI](https://github.com/create-go-app/fiber-go-template) ^^

## Setup
- To run development mode, execute `make watch`
- [Linux] Install `swag` by execute `go install github.com/swaggo/swag/cmd/swag@latest`
- Generate swagger docs by execute `make swagger`

## Features
- Login
- Register
- Verify Registration
- Reset Password
- My Profile
- Change Profile
- Example of CRUD function
- Verification on email change
- Role Permission Based Access Control inspired by [Spatie Laravel Permission Matrix](https://github.com/spatie/laravel-permission)

## Capabilities
- Auto migration
- Auto seeder

## Security Rules
- In this boilerplate, we use stateless JWT
- To change your email address, you need to verify the changes via email

## Screenshots
![image](https://user-images.githubusercontent.com/26473549/213926419-3d2d4b53-3060-48e6-9c26-ebb86c0466b2.png)
