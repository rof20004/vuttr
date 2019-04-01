FORMAT: 1A

# VUTTR

VUTTR is a simple API to manage and visualize very useful tools to remember.

# Group Tools

Resource related to tools in the API.

## Tool Collection [/tools]

### List All Tools [GET]

+ Response 200 (application/json)

        [
            {
              id: 1,
              title: "Notion",
              link: "https://notion.so",
              description: "All in one tool to organize teams and ideas. Write, plan, collaborate, and get organized. ",
              tags: [
                  "organization",
                  "planning",
                  "collaboration",
                  "writing",
                  "calendar"
              ]
            }
        ]

### Create a new Tool [POST]

Create a new tool. It takes a JSON object.

+ title (string) - The title of tool
+ link (string) - The URI of the tool
+ description (string) - A description about the tool
+ tags [array(string)] - A collection of identifiers of the tool

+ Request 200 (application/json)

        {
            "title": "hotel",
            "link": "https://github.com/typicode/hotel",
            "description": "Local app manager. Start apps within your browser, developer tool with local .localhost domain and https out of the box.",
            "tags":["node", "organizing", "webapps", "domain", "developer", "https", "proxy"]
        }

+ Response 201 (application/json)

  + Headers

            Location: /tools/5

  + Body

            {
                "title": "hotel",
                "link": "https://github.com/typicode/hotel",
                "description": "Local app manager. Start apps within your browser, developer tool with local .localhost domain and https out of the box.",
                "tags":["node", "organizing", "webapps", "domain", "developer", "https", "proxy"],
                "id":5
            }

## Tool [/tools/{tool_id}]

+ Parameters
  + tool_id (number) - ID of the Tool in the form of an integer

### View a Tool in Detail [GET]

Get details of a tool

+ Response 200 (application/json)

        {
            id: 1,
            title: "Notion",
            link: "https://notion.so",
            description: "All in one tool to organize teams and ideas. Write, plan, collaborate, and get organized. ",
            tags: [
                "organization",
                "planning",
                "collaboration",
                "writing",
                "calendar"
            ]
        }

### Delete [DELETE]

Delete a tool from registry

+ Response 204

## Search Tool by Tag [/tool?tag={tag_name}]

+ Parameters
  + tag_name (string) - A tag identifier of the tool

### Retrieve a list of tools by tag name [GET]

+ Response 200 (application/json)

        [
            {
                id: 2,
                title: "json-server",
                link: "https://github.com/typicode/json-server",
                description: "Fake REST API based on a json schema. Useful for mocking and creating APIs for front-end devs to consume in coding challenges.",
                tags: [
                    "api",
                    "json",
                    "schema",
                    "node",
                    "github",
                    "rest"
                ]
            }
        ]