package main

import (
  "log"
  "net"
  "net/http"
  "net/rpc"
)

type HouseDAO struct {
  Id         int
  Name       string
  Region     string
  CoatOfArms string
  Words      string
}

var houses = []HouseDAO{
  HouseDAO{
    Id:         1,
    Name:       "House Algood",
    Region:     "The Westerlands",
    CoatOfArms: "A golden wreath, on a blue field with a gold border(Azure, a garland of laurel within a bordure or)",
    Words:      "",
  },
  HouseDAO{
    Id:         2,
    Name:       "House Allyrion of Godsgrace",
    Region:     "Dorne",
    CoatOfArms: "Gyronny Gules and Sable, a hand couped Or",
    Words:      "No Foe May Pass",
  },
  HouseDAO{
    Id:         3,
    Name:       "House Amber",
    Region:     "The North",
    CoatOfArms: "",
    Words:      "",
  },
}

type Args struct {
  Id int
}

type House int

func (house *House) GetHouse(Id int, reply *HouseDAO) error {
  var found HouseDAO;

  for _, v := range houses {
    if v.Id == Id {
      found = v;
    }
  }
  *reply = found;
  return nil;
}

func (house *House) GetHouses(Id int, reply *[]HouseDAO) error {
  *reply = houses;
  return nil;
}

func main() {
  house := new(House);
  // // rpc.Register(house);
  // rpc.Register(house);
  
  // rpc.HandleHTTP();
  listener, e := net.Listen("tcp", "localhost:1234");
  if e != nil {
    log.Fatal("Listen error : ", e);
  }
  log.Printf("Starting server on port 1234");

  server := rpc.NewServer();
  server.Register(house);

  http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")
    server.ServeHTTP(w, r);
  })
  err := http.Serve(listener, nil);

  if err != nil {
    log.Fatal("Error serving : ", err);
  }
}
