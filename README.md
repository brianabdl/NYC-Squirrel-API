# NYC Squirrel API
This is a simple API that returns a list of squirrels in New York City. It is a RESTful API that allows you to get a list of all squirrels, get a specific squirrel by ID, filter it.


## Background
This API is based on the 2018 New York City Squirrel Census data. The data is available in CSV format and contains information about squirrels in New York City, including their ID, age, primary fur color, highlight fur color, and more.
I created this API to practice my Go programming skills and learn how to build a simple RESTful API.

## Endpoints
- GET /
- GET /ping
- GET /id/:id
- GET /age/:age
- GET /primaryFurColor/:color
- GET /highlightFurColor/:color
- GET /condition/:cond

## Usage
To use this API, simply make a GET request to the desired endpoint. For example, to get a list of all squirrels, make a GET request to `/`. To get a specific squirrel by ID, make a GET request to `/id/:id`, where `:id` is the ID of the squirrel you want to get.

## Example
Here is an example of how to use this API to get a list of all squirrels:

```bash
curl http://localhost:8080/
```

This will return a JSON response with a list of all squirrels in New York City Central Park.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
MIT License

Copyright (c) 2024 Muhammad Brian Abdillah
```