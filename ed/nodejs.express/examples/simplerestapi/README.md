Simple REST API
-

[<img src="https://heroku-badge.herokuapp.com/?app=heroku-badge">](https://simplerestapiforgta.herokuapp.com/)

This is simple REST API, which provides access to next resources:

* users
* customers
* countries

## Usage

rest API supports next methods:

- GET api/users, api/users/:id
- GET api/customers, api/customers/:id
- GET api/countries, api/countries/:id

For example:

````bash
curl https://simplerestapiforgta.herokuapp.com/users
curl https://simplerestapiforgta.herokuapp.com/users/1
````

Feel free do anything you wish! 🤓
in case of error you will see something like this:

````json
{"errors":[{"param":"id","msg":"Invalid id value.","value":"1x"}]}
````

## Most interesting

This project uses npm package [getthemall](https://github.com/cn007b/getthemall)
which provides ability to fetch data from different REST API endpoints into one request, like:

````bash
curl 'https://simplerestapiforgta.herokuapp.com/resources?user=users/1&my_customer=customers/2&countries=countries'
````

as result you will see:
````json
{
  "countries": [
    {
      "_id": "594afc85734d1d28ba613895",
      "code": "GB",
      "id": 1,
      "name": "United Kingdom"
    },
    {
      "_id": "594afcb8734d1d28ba61389c",
      "code": "US",
      "id": 3,
      "name": "United States"
    },
    {
      "_id": "594afca8734d1d28ba613898",
      "code": "UA",
      "id": 2,
      "name": "Ukraine"
    }
  ],
  "my_customer": {
    "_id": "594afe63734d1d28ba61398c",
    "date": "12/5/2016",
    "email": "cshiels1@studiopress.com",
    "full_name": "Colin Shiels",
    "id": 2,
    "ip_address": "90.5.250.164",
    "phone": "45-(832)802-1627"
  },
  "user": {
    "_id": "594afd91734d1d28ba613906",
    "email": "stenant0@seattletimes.com",
    "first_name": "Shelden",
    "gender": "Male",
    "id": 1,
    "last_name": "Tenant"
  }
}
````
