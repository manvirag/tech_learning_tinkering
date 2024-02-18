## Design the Google map.

### Functional Requirements

1. Support User location update. ( live location analysis and update in map)
2. navigation service including the ETA. ( find the optimal route, eta from source to destination by distance and time, depending on the type of transportation)
3. Map rendering ( find the particular place )

### Non-functional Requirements

1. Accuracy.
2. General availability and scalability

### Pre-requisite

#### Positioning system
For a particular point in the earth, there is two variable. i.e. 
![alt_text](./images/img.png)
- Latitute (Lat):    denotes how far north or south we are. 
- Longitude (Long):  denotes how far east or west we are.

#### Going from 3D to 2D

- This process is called **Map Projection**. ( Google Use Web Meracator)

#### Geocoding
- Process to find the lat and long from the address or name.
- Sample E.g.  Manit -> (23,45).

#### Geohashing
![alt_text](./images/img_1.png)
- Encode the geographic area into digits and letters.
- Its is recursive mean, on zooming one tile it generate its 4*x.

#### Map Rendering
![alt_text](./images/img_2.png)
![alt_text](./images/img_3.png)
- Tiling -> We show the tile on the zoom level.
- Let say global there is one tile, if someone zoom it then it shows the tiles which are inside the above zoom level.

#### Road data processing for navigation algorithm

![alt_text](./images/img_4.png)
https://medium.com/@sidgangs99/developing-google-maps-3c3320a365d

- For finding the optimal path algorithm A* or dijkstra variations can be used.
- Important part is how are we doing. Since these works on graph i.e. node and edges. Node denoting the location and edges denoting road.
- Denoting whole world like this would be inefficient take too memory and wouldn't be good for these algorithm.
- That's why we have Routing tiles. Same like Map rendering tiles. It is also mantained on the zoom level. Each zoom level have differnt route. like More upper only connecting highways or state, then district ...etc
- These tile have its graph inside. and these tile connected to zoom level.
- Means images like zoom 1 tile -> itself a graph.
- and zoom 2 which is made of zoom2 . That's how reach from zoom1 graph to high level.

### Capacity Estimation

- Storage: this can be calculated with total no. of tiles ( already shown in above pic ) and its image size.
- Optimization: 80-90% area are natural these images can be compressed . So can consider only 30% of total.
- Assume 1 billions DAU use 30 minutes map per week. or 5 billion minutes / day 
- Assume 1 request per second for live location from client.
- then 5*60 -> 300 billions request per day.
- 3 million / second.
- But we can optimize it since its not compulsory to send live location per second , we can batch on client side and sent. Assume 15 second perrequest.
- 200000 request/second -> PQPS *5 -> 1 million QPS.

### High level design
![alt_text](./images/img_5.png)

1. Location Service:
- That get the user live location and store in db for the future reference and usecase . ( Useful in analysis and get to know user live location )
- We the user location in batch from the client.
- Database: High write and scalable -> nosql or coloum  -> cassandra.
- Protocol: Http POST request.

2. Navigation Service:
- Take the source and destination location and return the routing data on that zoom level. (Reroute and ETA in deep dive.)

3. Map Rendering:

![alt_text](./images/img_6.png)
- For showing the data on the particular zoom level, we send the map tiles.
- Option 1: Server tell on the fly. Bad data is huge and static no caching .
- Option 2: CDN , we already send the pre generated tiles to the CDN and then client use it. Since each map tile have its unique string.
- For getting the url of the CDN for map tiles , we can introduce another service , that will take the lat and long with the zoom level and tell the map render tiles and also the other surrounding 8 tiles.
- Then client with these url can get the tiles.

### Deep dive in

#### References:

1. Alex xu system design vol 2
2. https://blog.bytebytego.com/p/design-google-maps