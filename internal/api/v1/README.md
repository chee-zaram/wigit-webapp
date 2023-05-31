# API V1

This is the first version of the backend API for the WIGIT Web Application.

The following are endpoints implemented and the data needed in the JSON payload
where applicable.

> NB: All routes should carry the prefix `/api/v1`.

> Strong fields are required.

> Routes that require sign in carry a star \*

> Admin routes carry two stars \*\*

> Super Admin routes carry three stars \*\*\*

## GET

<ul>
    <li><h4>/products</h4>
    Get all products in the database.
    On success, it will return a <b>200</b> response code with a <b>data</b> object which is a list of all product objects in the payload.
    On failure, it will return a <b>500</b> response code and an <b>error</b> string in the payload.
    </li>
    <li><h4>/products/:product_id</h4>
    Get a particular product from the database.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code and an <b>error</b> string in the payload.
    </li>
    <li><h4>/products/categories/:category</h4>
    Get all products in a given category.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload. This will be a list of all product objects in
    the given category.
    On failure, it will return a <b>400</b> or <b>500</b> response code and an <b>error</b> string in the payload.
    </li>
    <li><h4>/cart *</h4>
    Get all items in the user's cart.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload. This will be a list of all item objects in
    the user's cart.
    On failure, it will return a <b>500</b> response code and an <b>error</b> string in the payload.
    </li>
    <li><h4>/services</h4>
    Get a list of all services.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload. This will be a list of all service objects.
    On failure, it will return a <b>500</b> response code and an <b>error</b> string in the payload.
    </li>
    <li><h4>/services/:service_id</h4>
    Get a particular service.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload which will be a service object.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/bookings *</h4>
    Get a list of the user's bookings.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload which will be a list of booking objects.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/bookings/:booking_id *</h4>
    Get a specific booking with booking_id belonging to the user.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload which will be a booking object.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/bookings **</h4>
    Get a list of all bookings in the database.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload which will be a list of booking objects.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/bookings/:booking_id **</h4>
    Get a specific booking with booking_id.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload which will be a booking object.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/slots</h4>
    Get all free slots which are still valid.
    On success, it will return a <b>200</b> response code with a <b>data</b> object in the payload which will be a slot object.
    On failure, it will return a <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
</ul>

## POST

<ul>
    <li><h4>/signup</h4>
    Sign a user up.
    <ul>
      <li><strong>email</strong>: The new user's email address. Must be unique and between 5 and 45 characters long.</li>
      <li><strong>password</strong>: the user's password. Between 8 and 45 characters.</li>
      <li><strong>repeat_password</strong>: A repeat of the user's password.</li>
      <li><strong>first_name</strong>: The user's first name. Not more than 45 characters.</li>
      <li><strong>last_name</strong>: The user's last name. Not more than 45 characters.</li>
      <li><strong>address</strong>: The user's address. Not more than 255 characters.</li>
      <li><strong>phone</strong>: The user's phone number. Between 9 and 11 characters.</li>
    </ul>
    On success, it will return a <b>201</b> response code with a <b>jwt</b> string and a <b>msg</b> string in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/signin</h4>
    Sign a user with given credentials in.
    <ul>
      <li><strong>email</strong>: The user's email address. Between 5 and 45 characters long.</li>
      <li><strong>password</strong>: The user's password. Between 8 and 45 characters.</li>
    </ul>
    On success, it will return a <b>200</b> response code with a <b>jwt</b> string, a <b>msg</b> string, and a <b>user</b> object in the payload.
    On failure, it will return a <b>400</b>, <b>401</b>, or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/reset_password</h4>
    Send a request to reset password for user with given email.
    <ul>
      <li><strong>email</strong>: The user's email address. Between 5 and 45 characters long.</li>
    </ul>
    On success, it will return a <b>201</b> response code with a <b>reset_token</b> string in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/products **</h4>
    Add a new product to the database.
    <ul>
      <li><strong>name</strong>: The name of the product. Must be 3 to 45 characters long.</li>
      <li><strong>description</strong>: The details of what the product is. Must be between 3 to 1024 characters long.</li>
      <li><strong>category</strong>: The category the product belong to. Must be 3 to 45 characters long.</li>
      <li><strong>stock</strong>: The quantity of the product in stock. An integer.</li>
      <li><strong>price</strong>: The price of the product. Decimal as a string. Will be rounded up to 2 decimal places.</li>
      <li><strong>image_url</strong>: A link to the product display image. Not longer than 255 characters.</li>
    </ul>
    On success, it will return a <b>201</b> response code with a <b>msg</b> string and a <b>data</b> object in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/cart *</h4>
    Add an item to the user's cart.
    <ul>
      <li><strong>product_id</strong>: This is the id of the product to add to cart.</li>
      <li><strong>quantity</strong>: This is the quantity of the product the user wants. An integer. Cannot be 0.</li>
    </ul>
    On success, it will return a <b>201</b> response code with a <b>msg</b> string and a <b>data</b> object in the payload. The data is the new item.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/services **</h4>
    Add a new service to the database.
    <ul>
      <li><strong>name</strong>: The name of the service. Must be 3 to 45 characters long.</li>
      <li><strong>description</strong>: The details of what the service is. Must be between 3 to 1024 characters long.</li>
      <li><strong>price</strong>: The price of the service. Decimal as a string. Will be rounded up to 2 decimal places.</li>
      <li><strong>available</strong>: A boolean. Says if the service is currently available or not.</li>
    </ul>
    On success, it will return a <b>201</b> response code with a <b>msg</b> string and a <b>data</b> object in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/bookings *</h4>
    Add a new user booking to the database.
    <ul>
      <li><strong>slot_id</strong>: The id for the slot the user has been booked for.</li>
      <li><strong>service_id</strong>: The id for the service the user has booked.</li>
    </ul>
    On success, it will return a <b>201</b> response code with a <b>msg</b> string and a <b>data</b> list of booking objects in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/slots **</h4>
    Add a new slot to the database.
    <ul>
      <li><strong>date_time</strong>: This is the date and time for the slot. String value format of a datetime. y/m/d "2006-01-02T15:00:00Z"</li>
      <li><strong>is_free</strong>: A boolean which says if the slot is free or not.</li>
    </ul>
    On success, it will return a <b>201</b> response code with a <b>data</b> object in the payload which will be a slot object.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
</ul>

## PUT

<ul>
    <li><h4>/reset_password</h4>
    <ul>
      <li><strong>email</strong>: The user's email address. Between 5 and 45 characters long.</li>
      <li><strong>reset_token</strong>: The reset token sent back for that user when a post to this route was made.</li>
      <li><strong>new_password</strong>: The user's new password. Must be 8 to 45 characters long.</li>
      <li><strong>repeat_new_password</strong>: A repeat of the user's new password.</li>
    </ul>
    On success, it will return a <b>200</b> response code with a <b>msg</b> string in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/products/:product_id **</h4>
    <ul>
      <li><strong>name</strong>: The name of the product. Must be 3 to 45 characters long.</li>
      <li><strong>description</strong>: The details of what the product is. Must be between 3 to 1024 characters long.</li>
      <li><strong>category</strong>: The category the product belong to. Must be 3 to 45 characters long.</li>
      <li><strong>stock</strong>: The quantity of the product in stock. An integer.</li>
      <li><strong>price</strong>: The price of the product. Decimal as a string. Will be rounded up to 2 decimal places.</li>
      <li><strong>image_url</strong>: A link to the product display image. Not longer than 128 characters.</li>
    </ul>
    On success, it will return a <b>200</b> response code with a <b>msg</b> string and a <b>data</b> object in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/services **</h4>
    <ul>
      <li><strong>name</strong>: The name of the service. Must be 3 to 45 characters long.</li>
      <li><strong>description</strong>: The details of what the service is. Must be between 3 to 1024 characters long.</li>
      <li><strong>price</strong>: The price of the service. Decimal as a string. Will be rounded up to 2 decimal places.</li>
      <li><strong>available</strong>: A boolean. Says if the service is currently available or not.</li>
    </ul>
    On success, it will return a <b>200</b> response code with a <b>msg</b> string and a <b>data</b> object in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/bookings/:booking_id/:status **</h4>
    This updates the status of a booking. Allowed status' are <b>pending</b>, <b>paid</b>, <b>fulfilled</b>, and <b>cancelled</b>.
    On success, it will return a <b>200</b> response code with a <b>msg</b> string and a <b>data</b> object of the booking in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
</ul>

## DELETE

<ul>
    <li><h4>/admin/products/:product_id **</h4>
    Delete a product with product_id.
    On success, it will return a <b>200</b> response code with a <b>msg</b> string and a <b>data</b> object in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/cart/:item_id *</h4>
    Delete an item with item_id from the user's cart.
    On success, it will return a <b>200</b> response code with a <b>msg</b> string in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/cart *</h4>
    Clear the user's cart.
    On success, it will return a <b>200</b> response code with a <b>msg</b> string in the payload.
    On failure, it will return a <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/services/:service_id **</h4>
    Delete a service from the database.
    On success, it will return a <b>200</b> response code with a <b>msg</b> string in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
    <li><h4>/admin/slots/:slot_id **</h4>
    Delete a slot from the database.
    On success, it will return a <b>200</b> response code with a <b>msg</b> string in the payload.
    On failure, it will return a <b>400</b> or <b>500</b> response code, and an <b>error</b> string in the payload.
    </li>
</ul>

<!-- TODO: add documentation for orders endpoint -->
<!-- TODO: add documentation for super_admin endpoint -->
<!-- TODO: add documentation for getting all user's orders and bookings -->

## Author

<details>
    <summary>Emmanuel Chee-zaram Okeke</summary>
    <ul>
    <li><a href="https://www.github.com/chee-zaram">GitHub</a></li>
    <li><a href="https://www.twitter.com/CheezaramOkeke">Twitter</a></li>
    <li><a href="https://www.linkedin.com/in/chee-zaram">Linkedin</a></li>
    <li><a href="mailto:ecokeke21@gmail.com">Gmail</a></li>
    </ul>
</details>
