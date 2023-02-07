![image](https://user-images.githubusercontent.com/26473549/214473471-19e5a263-cf21-440a-9beb-d2cc16eed9fc.png)

# A production ready Golang Boilerplate API


## Setup
- Copy `.env.example` to `.env`, adjust the values, and copy the `.env` to `test/unit` and `test/api`
- To run API test, execute `make api_test`
- To run unit test, execute `make unit_test`
- To run development mode, execute `make watch`
- [Linux] Install `swag` by execute `go install github.com/swaggo/swag/cmd/swag@latest`
- Generate swagger docs by execute `make swagger`

## Features
- Log in using username and password
- Log in using Google One Tap, please see the example in the folder `public/test_google_one_tap.html` to use Google One Tap in the front end, and endpoint `/api/v1/auth/google-one-tap` for the backend 
- Register
- Verify Registration
- Reset Password
- My Profile
- Change Profile
- Example of CRUD function
- Verification on email change
- Role Permission Based Access Control inspired by [Spatie Laravel Permission Matrix](https://github.com/spatie/laravel-permission)
- Pagination
- Cronjob using [gocron](https://github.com/go-co-op/gocron)

## Capabilities
- Auto migration
- Auto seeder

## Security Rules
- In this boilerplate, we use stateless JWT
- Send email notification when the user's profile is updated
- To change your email address, you need to verify the change via email
- Implementation of rate limiter in auth routes
- Implementation of mutex

## Screenshots
![image](https://user-images.githubusercontent.com/26473549/215829825-b6964b0f-ff95-4b4f-8ba8-e934758fbaa0.png)
![image](https://user-images.githubusercontent.com/26473549/213926419-3d2d4b53-3060-48e6-9c26-ebb86c0466b2.png)

## Install Gosec
- Run `go install github.com/securego/gosec/v2/cmd/gosec@latest`

## References
- [https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34](https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34)
- [Email template](https://codepen.io/mightyteja/pen/xxxjXqJ)
- Inspired by [Fiber backend template for Create Go App CLI](https://github.com/create-go-app/fiber-go-template) ^^
