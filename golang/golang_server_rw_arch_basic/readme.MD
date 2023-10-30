Doc: https://go.dev/doc/tutorial/web-service-gin
Backend Golang repository steps:

1. To keep things simple for the tutorial, you’ll store data in memory. A more typical API would interact with a database.
2. Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.
3. Add the function getAlbums
4. Add the code in main for routing and map route to function
5. import  "github.com/gin-gonic/gin"
6. Add more route and its handler