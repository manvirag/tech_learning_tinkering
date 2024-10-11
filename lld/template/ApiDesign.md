Template:
-  This is most simplist part.
-  According to requirements, we have to make the api 
    Restful Apis:

1. Decide → GET, POST ( body ) , PUT( body ), DELETE.
2. Route Naming: 
    1. Keep route simple.
        1. /api/v1/users
        2. /api/v1/users/{id}
        3. /api/v1/users?abc=’’&&abc=’’ → for more complex or filter or sorting.
    2. Versioning →  routes, like `/api/v1/..`
    3. Define input json → can require metadata.
    4. in case of pagination
        1. totalcount , pagesize, has next, hits .
    5. valid response → error output  with status code
        1. 500
        2. 200
        3. 400
        4. 404
        5. proper error message.
