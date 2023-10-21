<div align="center">
  <a href="https://webapp-wigit-ng.vercel.app/"><img alt="Wigit" src="https://github.com/wigit-gh/.github/raw/main/wigit.png" width="558" /></a>
  <br/>
  <strong>Bringing wig products and services online for easy access and convenience in Ghana</strong>
  <h1>WIG!T Web Application</h1>
</div>

[![Workflow](https://github.com/wigit-ng/webapp/actions/workflows/backend.yml/badge.svg)][workflow]
[![Go Report](https://goreportcard.com/badge/github.com/wigit-ng/webapp/backend)][report]
![Last Commit](https://img.shields.io/github/last-commit/wigit-ng/webapp)
![Contributors](https://img.shields.io/github/contributors/wigit-ng/webapp)

---

## Table of Contents

- [Introduction](#introduction)
- [Starting the Application](#starting-the-application)
  - [Backend](#backend)
  - [Frontend](#frontend)
- [Usage](#usage)
- [Developers](#developers)
  - [Backend](#dev_backend)
  - [Frontend](#dev_frontend)
  - [DevOps](#dev_devops)
- [Contributing](#contributing)
- [Related Projects](#related-projects)
- [Licensing](#licensing)

## Introduction

The **WIG!T Web Application** is a full-stack web project that aims to restore
customer trust in online shopping. Our mission is to bring shopping closer to
the people. We are inspired by the need to provide safety, assurance, quality,
and peace of mind to customers who simply want to experience the joy and
convenience of shopping.

The Web Application aims to reduce customer wait times by providing seamless
user experience, customer service, and delivery experience. We keep customers
informed and reassured every step of the way.

To learn more about the **WIG!T** brand, you can:

- Visit the [landing page](https://webapp-wigit-ng.vercel.app/).
- Know the [founder](https://webapp-wigit-ng.vercel.app/) and
  [developers](#developers).
- Read our [blog post](https://webapp-wigit-ng.vercel.app/) on the launch of the
  Web Application.
- Try out [our application](https://webapp-wigit-ng.vercel.app/).

## Starting the Application

The procedure outlined expects that you're setting up for testing or
development. The DevOps team will be responsible for server setup and
configuration required in a production environment.

Provided you use **Linux/GNU** and have **git** installed, the application can
be started by cloning this repository on the command line using the following
command:

```sh
git clone https://github.com/wigit-ng/webapp.git wigit-webapp
```

changing working directory into the application directory:

```sh
cd wigit-webapp
```

### [Backend](/backend)

The backend is written in [Go Programming Language](https://go.dev/) and uses
the [Gin Web Framework](https://gin-gonic.com/). Server configurations will be
carried out by the DevOps team in production prior to backend deployment. The
following instructions apply to devolopment and testing. Documentation for the
backend API is available on
[GitHub](https://github.com/wigit-ng/webapp/blob/main/backend/internal/api/v1/README.md).
Documentation on the back-end API has also been done using
[Swagger](https://backend.wigit.com.ng/api/v1/swagger/index.html).

##### Dependencies

You only need [docker](https://docs.docker.com/engine/install/) installed.

Navigate to the [backend](/backend) directory from the root of the project:

```sh
cd backend
```

Run the docker compose command:

```sh
docker compose up --build -d
```

to create three required services, and start three required containers namely:

- wigit
- wigit-mysql
- wigit-redis

Run `docker compose ps` to see the containers running.

The `wigit` container runs the main app, and listens for connections on port
`8080`.

> NB: You can edit this
> [sample.env](https://github/wigit-ng/webapp/backend/sample.env) file based on
> your preferred backend configuration, and then rename it to `.env` to allow
> docker compose use it instead of the defaults, if you wish.

You can now send requests to the backend API on port `8080`. You can use `cURL`,
E.g but I like to use [xh](https://github.com/ducaale/xh). Cleaner output and
easier syntax.

xh:

```sh
xh :8080/api/v1/products
```

cURL:

```sh
curl http://localhost:8080/api/v1/products
```

You will need to create a regular user account:

xh:

```sh
xh post :8080/api/v1/signup email='yours@email.com' password='password' \
repeat_password='password' first_name='John' last_name='Doe' \
address='No 10. Nothing Rd' phone=07038639012
```

cURL:

```sh
curl -X POST http://localhost:8080/api/v1/signup \
-H "Content-Type: application/json" \
-d '{
  "email": "yours@email.com",
  "password": "password",
  "repeat_password": "password",
  "first_name": "John",
  "last_name": "Doe",
  "address": "No 10. Nothing Rd",
  "phone": "07038639012"
}'
```

Then, update the role from **customer** to **admin** to enable you manage
products and services by starting up a shell in the `wigit-mysql` container:

```sh
docker exec -it wigit-mysql bash
```

Go into the database:

```sh
mysql
```

Update the newly created user role to `admin`:

```sh
UPDATE wigit.users SET role = 'admin' WHERE email = 'yours@email.com'; exit;
```

Exit the `wigit-mysql` terminal by typing `exit`.

The above actions now gives the user permission to perform admin duties, like
managing products and services.

JWT is used for authentication. A sample request to add a new product looks
like:

xh:

```sh
xh post :8080/api/v1/admin/products Authorization:'Bearer token_returned_on_signup' \
name='Ghanaian wig' description='A custom made wig for ghana' stock:=10 price:=300 \
image_url='https://images.pexels.com/photos/973406/pexels-photo-973406.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1' \
category='straight'
```

cURL:

```sh
curl -X POST http://localhost:8080/api/v1/admin/products \
-H "Authorization: Bearer token_returned_on_signup" \
-H "Content-Type: application/json" \
-d '{
  "name": "Ghanaian wig",
  "description": "A custom made wig for Ghana",
  "stock": 10,
  "price": 300,
  "image_url": "https://images.pexels.com/photos/973406/pexels-photo-973406.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
  "category": "straight"
}'
```

The docker environment is setup to persist the data between runs, or even if a
new container is spun, as volumes are mounted in the `wigit-docker-data`
directory for the database and logs.

Visit [the docs](/backend/internal/api/v1/README.md) to see more available
endpoints and required fields.

### [Frontend](/frontend)

For the frontend, the [initial steps](#starting-the-application) are carried out
on the already configured frontend server in production, or anywhere for testing
and devolopment purposes. The frontend depends on a deployed backend whose
hostname is used in the source code.

##### Dependencies

- [Node and npm](https://github.com/nvm-sh/nvm)

Next, change directory into the [frontend](/frontend) directory which is where
the frontend source code resides and is the root of the frontend built with
[Next.js](https://nextjs.org/).

```sh
cd frontend
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
the deployed application at
[the webapp](https://webapp-wigit-ng.vercel.app/contact).

Anyone can perform the following:

- Visit the homepage
- View `Products`
- View `Services`
- View `About` page
- View `Contact Us` page

Signed-in users can additionally:

- Add items to the cart
- View cart
- Place orders
- Book services
- Track orders and bookings
- View and edit profile
- Delete their account

Admins can:

- View customer orders and bookings
- Change orders and bookings status

Super Admins can:

- Get all user information
- See all admins
- See all customers
- Update previleges for a user account
- Delete an account

## Developers

This lists all individuals who have contributed to the development of this
application. Their full names, links, and contact information are listed below:

<div id="dev_backend">
  <h4>Backend</h4>
  <details>
    <summary>Emmanuel Chee-zaram Okeke</summary>
    <ul>
    <li><a href="https://www.cheezaram.tech">Website</a></li>
    <li><a href="https://www.github.com/chee-zaram">GitHub</a></li>
    <li><a href="https://www.twitter.com/CheezaramOkeke">Twitter</a></li>
    <li><a href="https://www.linkedin.com/in/chee-zaram">Linkedin</a></li>
    <li><a href="mailto:ecokeke21@gmail.com">Gmail</a></li>
    </ul>
  </details>
</div>

---

<div id="dev_frontend">
  <h4>Frontend</h4>
  <details>
    <summary>Ovy Evbodi</summary>
    <ul>
    <li><a href="https://www.github.com/OvyEvbodi">GitHub</a></li>
    <li><a href="https://www.linkedin.com/in/ovy-evbodi-21920a203/">Linkedin</a></li>
    <li><a href="mailto:evbodiovo@gmail.com">Gmail</a></li>
    </ul>
  </details>
</div>

---

<div id="dev_devops">
  <h4>DevOps</h4>
  <details>
    <summary>Emmanuel Chee-zaram Okeke</summary>
    <ul>
    <li><a href="https://www.cheezaram.tech">Website</a></li>
    <li><a href="https://www.github.com/chee-zaram">GitHub</a></li>
    <li><a href="https://www.twitter.com/CheezaramOkeke">Twitter</a></li>
    <li><a href="https://www.linkedin.com/in/chee-zaram">Linkedin</a></li>
    <li><a href="mailto:ecokeke21@gmail.com">Gmail</a></li>
    </ul>
  </details>
</div>

## Contributing

Only members of the software engineering team can contribute to the source code.
To report bugs and issues, or make feature requests, kindly send us a mail
through our [support page](https://webapp-wigit-ng.vercel.app/contact) or
directly at our support email support@wigit.com.ng.

## Related Projects

Some project similar to ours include:

- wigwholesale.com
- wigsbyvanity.com

## Licensing

Copyright (c) 2023, WIG!T. All Rights Reserved

The code in this repository is not provided under an open source license. It is
solely intended for internal use within **WIG!T**. If you wish to use this code
or incorporate it into your own projects, please contact support@wigit.com.ng to
discuss licensing and obtain permission.

[workflow]: https://github.com/wigit-ng/webapp/actions/workflows/backend.yml?query=branch%3Amain+event%3Apush
[report]: https://goreportcard.com/report/github.com/wigit-ng/webapp/backend
