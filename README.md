# Project VUTTR

## Tools

1. Go
2. Docker
3. Docker Compose
4. API Blueprint
5. Fasthttp framework
6. PostgreSQL

## How-to

* Run project
  > make start
  
* Execute migrations
  > make migrations

## Endpoints

| URI                                     |  Method     |  Protected  |    Short Description   |
|-----------------------------------------|-------------|-------------|------------------------|
| http://localhost:3000/tools             |   GET       |             |  Query all tools       |
| http://localhost:3000/tools             |   POST      |             |  Create a tool         |
| http://localhost:3000/tools?tag=node    |   GET       |             |  Query tools by tag    |
| http://localhost:3000/tools/:tool_id    |   DELETE    |       YES   |  Delete tool by id     |


> Basic Auth: username = admin password = bossabox
  
 
### Enjoy!