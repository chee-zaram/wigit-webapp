# WIG!T Web Application

> Bringing wig products and services online for easy access and convenience in
> Ghana.

> ![Home Top](https://imgur.com/AMIfkWe.png)

> ![Home Trending](https://imgur.com/lAKPnhZ.png)

> ![Products Top](https://imgur.com/h5ALl13.png)

> ![Products Bottom](https://imgur.com/4OCT8Ea.png)

> ![Sign In](https://imgur.com/UddM0co.png)

## Table of Content

- [Introduction](#introduction)
- [Starting the Application](#starting-the-application)
  - [Backend](#backend)
  - [Frontend](#frontend)
- [Usage](#usage)
- [Contributing](#contributing)
- [Related Projects](#related-projects)
- [Licensing](#licensing)

## Introduction

The **WIG!T Web Application** is a full stack web project which aims to restore
customer trust in online shopping. Our mission is to bring shopping closer to
the people. We are inspired by a need to provide safety, assurance, quality, and
peace-of-mind to customers who simply want to experience the joy and convenience
of shopping.

The Web Application aims to reduce customer wait times by providing seemless
user experience, customer service, and delivery experience. We keep customers
informed and reassured every step of the way.

To learn more about the **WIG!T** brand, you can:

- Visit the [landing page]().
- Know the [founder]() and [engineers]().
- Read our [blog post]() on the launch of the Web Application.
- Try out [our application]().

## Starting the Application

The procedure outlined expect you're setting up for testing or devolopment. The
DevOps team will be responsible for server setup and configuration required in a
production environment.

Provided you use **Linux/GNU** and have **git** installed, the application can
be started by cloning this repository on the commandline using the following
command:

```sh
git clone https://github.com/wigit-gh/webapp.git wigit-webapp
```

changing working directory into the application directory:

```sh
cd wigit-webapp
```

### Backend

The backend is written in [Go Programming Language](https://go.dev/) and uses
the [Gin Web Framework](https://gin-gonic.com/). Server configurations will be
carried out by the DevOps team in production prior to backend deployment. The
following instructions apply to devolopment and testing. Documentation for the
backend API is available on
[GitHub](https://github.com/wigit-gh/webapp/blob/main/internal/api/v1/README.md).
Documentation on the API has also been done using
[Swagger](https://cheezaram.tech/api/v1/swagger/index.json).

##### Dependencies

- Go 1.2.x
- MySQL 8.x

After carrying out the [initial steps](#starting-the-application) and setting up
dependencies, install all required modules:

```sh
go mod tidy
```

start the backend using:

```sh
go run main.go
```

This will start the backend server to listen on all hosts on port `8000`.
**Gin** will also start in debug mode which should make all routes visible on
start-up.

### Frontend

For the frontend, the [initial steps](#starting-the-application) are carried out
on the already configured frontend server in production, or anywhere for testing
and devolopment purposes. The frontend depends on a deployed backend whose host
name is used in the source code.

##### Dependencies

- [Node and npm](https://github.com/nvm-sh/nvm)

Next, change directory into the
[assets](https://github.com/wigit-gh/webapp/tree/main/assets) directory which is
where the frontend source code resides and is the root of the frontend built
with [Next.js](https://nextjs.org/).

```sh
cd assets
```

install all dependencies needed with:

```sh
npm install
```

In production, the Frontend is built using `npm build`, and then the static
files are duly deployed on the server. For testing purposes, the devolopment
server will do. Start it by running:

```sh
npm run dev
```

Visit `http://localhost:3000` on your browser to interact with the application.

## Usage

The fullstack application can be interacted with from the browser by visiting
the deployed application at [the website]().

Anyone can perform the following:

- Visit the homepage
- View `Products`
- View `Services`
- View `About` page
- View `Contact Us` page

Signed in users can additionally:

- Add items to cart
- View cart
- Place orders
- Book services
- Track orders and services
- View and edit profile

Admins can:

- View customer orders and bookings
- Change orders and bookings status

Super Admins can:

- Get all users information
- See all admins
- See all customers
- Delete a user account
- Update previleges for a user account

## Contributing

Only members of the software engineering team can contribute to the source code.
To report bugs and issues, or make feature requests, kindly send us a mail
through our [support page]() or directly at our support email support@wigit.com.

## Related Projects

Some project similar to ours include:

- wigwholesale.com
- wigsbyvanity.com

## Licensing

The code in this repository is not provided under an open source license. It is
solely intended for internal use within **WIG!T**. If you wish to use this code
or incorporate it into your own projects, please contact support@wigit.com to
discuss licensing and obtain permission.
