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
- Upload file to AWS S3

## Middlewares
### Using rate limiter
We have 3 rate limiters here:
- ThrottleByIp, this rate limiter will limits requests by the user's IP, if you are using Nginx, please make sure that the x-real-ip header is exposed to the back end
- ThrottleByKeyAndIP, this rate limiter will limit requests by the user's IP and a custom key, for example, if you need to rate limit an endpoint to register, you may use this rate limiter by ThrottleByKeyAndIP("register", 60, 60). It will limits request from an IP with the key "register" to only 60 times each 60 seconds
- ThrottleByKey, this rate limiter will limit requests by a custom key, for example, if you need to rate limit an endpoint to register, you may use this rate limiter by ThrottleByKey("register", 60, 60). It will limits request from any IP with the key "register" to only 60 times each 60 seconds
For the real example, you may open `app/routes/api/v1/auth.go`

### Protect a route using JWT
You can use `PermissionAuth` middleware to protect a URL by the current user's permissions. It means that the users must have particular permission(s) in order to access the protected routes.
You need to import `goshaka/app/middlewares` and then add the `middlewares.PermissionAuth` middleware as a route handler, 
you can see the example in `app/routes/api/v1/user.go`.
For example `middlewares.PermissionAuth([]string{"user-read"})` , it means in order to access a particular route, a user must have permission `user-read`

### Protect a route using user's permission checker 
You can use `PermissionAuth` middleware to protect a URL by the current user's permissions. It means that the users must have particular permission(s) in order to access the protected routes.
You need to import `goshaka/app/middlewares` and then add the `middlewares.PermissionAuth` middleware as a route handler, 
you can see the example in `app/routes/api/v1/user.go`.
For example `middlewares.PermissionAuth([]string{"user-read"})` , it means in order to access particular route, a user must have permission `user-read`

### Protect a route using user's role checker 
You can use `RoleAuth` middleware to protect a URL from the current user's roles. It means that the users must have a particular role(s) in order to access the protected routes.
You need to import `goshaka/app/middlewares` and then add the `middlewares.RoleAuth` middleware as a route handler, 
you can see the example in `app/routes/api/v1/user.go`.
For example `middlewares.RoleAuth([]string{"admin"})` , it means in order to access the particular route, a user must have the role `admin`

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
### Email
![image](https://user-images.githubusercontent.com/26473549/215829825-b6964b0f-ff95-4b4f-8ba8-e934758fbaa0.png)
### Swagger
![image](https://user-images.githubusercontent.com/26473549/213926419-3d2d4b53-3060-48e6-9c26-ebb86c0466b2.png)

## Install Gosec
- Run `go install github.com/securego/gosec/v2/cmd/gosec@latest`

## References
- [https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34](https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34)
- [Email template](https://codepen.io/mightyteja/pen/xxxjXqJ)
- Inspired by [Fiber backend template for Create Go App CLI](https://github.com/create-go-app/fiber-go-template) ^^
- https://medium.com/spankie/upload-images-to-aws-s3-bucket-in-a-golang-web-application-2612bea70dd8
