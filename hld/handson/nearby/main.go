package main

import (
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Location represents a point with its geohash
type Location struct {
	ID        int64
	Name      string
	Latitude  float64
	Longitude float64
	GeoHash   string
	CreatedAt time.Time
}

const base32 = "0123456789bcdefghjkmnpqrstuvwxyz"

// getGeoHash converts latitude and longitude to a geohash of specified precision
func getGeoHash(lat, lon float64, precision int) string {
	var geohash string
	var minLat, maxLat float64 = -90, 90
	var minLon, maxLon float64 = -180, 180
	var mid float64
	var bit, ch int
	even := true

	for len(geohash) < precision {
		if even {
			mid = (minLon + maxLon) / 2
			if lon > mid {
				ch |= 1 << (4 - bit)
				minLon = mid
			} else {
				maxLon = mid
			}
		} else {
			mid = (minLat + maxLat) / 2
			if lat > mid {
				ch |= 1 << (4 - bit)
				minLat = mid
			} else {
				maxLat = mid
			}
		}
		even = !even

		if bit < 4 {
			bit++
		} else {
			geohash += string(base32[ch])
			bit = 0
			ch = 0
		}
	}
	return geohash
}

// decodeBoundingBox returns the bounding box for a geohash
func decodeBoundingBox(geohash string) (minLat, maxLat, minLon, maxLon float64) {
	minLat, maxLat = -90.0, 90.0
	minLon, maxLon = -180.0, 180.0
	even := true

	for _, c := range geohash {
		idx := strings.IndexRune(base32, c)
		if idx == -1 {
			panic("invalid geohash character")
		}
		for mask := 16; mask > 0; mask >>= 1 {
			if even {
				mid := (minLon + maxLon) / 2
				if idx&mask != 0 {
					minLon = mid
				} else {
					maxLon = mid
				}
			} else {
				mid := (minLat + maxLat) / 2
				if idx&mask != 0 {
					minLat = mid
				} else {
					maxLat = mid
				}
			}
			even = !even
		}
	}
	return
}

// getNeighbors returns 8 neighboring geohashes around the given geohash
func getNeighbors(hash string) []string {
	minLat, maxLat, minLon, maxLon := decodeBoundingBox(hash)
	latMid := (minLat + maxLat) / 2
	lonMid := (minLon + maxLon) / 2
	deltaLat := (maxLat - minLat)
	deltaLon := (maxLon - minLon)
	precision := len(hash)

	points := []struct{ lat, lon float64 }{
		{latMid + deltaLat, lonMid},            // N
		{latMid - deltaLat, lonMid},            // S
		{latMid, lonMid + deltaLon},            // E
		{latMid, lonMid - deltaLon},            // W
		{latMid + deltaLat, lonMid + deltaLon}, // NE
		{latMid + deltaLat, lonMid - deltaLon}, // NW
		{latMid - deltaLat, lonMid + deltaLon}, // SE
		{latMid - deltaLat, lonMid - deltaLon}, // SW
	}

	neighbors := make([]string, 0, 8)
	for _, p := range points {
		neighbors = append(neighbors, getGeoHash(p.lat, p.lon, precision))
	}
	return neighbors
}

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3 // Earth's radius in meters
	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) + math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

func main() {
	db, err := sql.Open("mysql", "nearby:nearby@tcp(localhost:3306)/nearby")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MySQL!")

	locations := []Location{
		{1, "Coffee Shop", 37.7749, -122.4194, "", time.Now()},
		{2, "Nearby Restaurant", 37.7750, -122.4195, "", time.Now()},
		{3, "Far Restaurant", 37.7850, -122.4294, "", time.Now()},
		{4, "Very Far Restaurant", 37.7950, -122.4394, "", time.Now()},
		{5, "North Edge", 37.7769, -122.4194, "", time.Now()},
		{6, "South Edge", 37.7729, -122.4194, "", time.Now()},
		{7, "East Edge", 37.7749, -122.4174, "", time.Now()},
		{8, "West Edge", 37.7749, -122.4214, "", time.Now()},
	}

	for i := range locations {
		locations[i].GeoHash = getGeoHash(locations[i].Latitude, locations[i].Longitude, 6)
		_, err := db.Exec("INSERT INTO locations (name, latitude, longitude, geohash) VALUES (?, ?, ?, ?)",
			locations[i].Name, locations[i].Latitude, locations[i].Longitude, locations[i].GeoHash)
		if err != nil {
			fmt.Printf("Error inserting location %s: %v\n", locations[i].Name, err)
			continue
		}
		fmt.Printf("Inserted location: %s\n", locations[i].Name)
	}

	coffeeShop := locations[0]
	prefixes := append([]string{coffeeShop.GeoHash[:4]}, getNeighbors(coffeeShop.GeoHash[:4])...)
	fmt.Printf("\nFinding locations near %s (geohash prefix: %s):\n", coffeeShop.Name, coffeeShop.GeoHash[:4])

	query := "SELECT id, name, latitude, longitude, geohash FROM locations WHERE "
	args := []interface{}{}
	for i, p := range prefixes {
		if i > 0 {
			query += " OR "
		}
		query += "geohash LIKE ?"
		args = append(args, p+"%")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var loc Location
		err := rows.Scan(&loc.ID, &loc.Name, &loc.Latitude, &loc.Longitude, &loc.GeoHash)
		if err != nil {
			panic(err)
		}

		if loc.Name == coffeeShop.Name {
			continue
		}

		distance := calculateDistance(coffeeShop.Latitude, coffeeShop.Longitude, loc.Latitude, loc.Longitude)
		fmt.Printf("%s is %.2fm away\n", loc.Name, distance)
	}
}
