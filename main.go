package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"

    "github.com/google/uuid"
)

var path_assets_img string = "assets/img/";
var output_path string = "/tmp/";

func main() {

    arg := os.Args[1]
    jsonFile, err := os.Open(arg)

    if isError(err) {
        return
    }

    defer jsonFile.Close()

    fmt.Println("Successfully Opened users.json")


    byteValue, _ := ioutil.ReadAll(jsonFile)

    var resources Resources

    json.Unmarshal(byteValue, &resources)

   out := Vpc{}
    for i := 0; i < len(resources.Resources); i++ {
        if len(resources.Resources[i].Instances) > 0 {
          if resources.Resources[i].Type == "aws_vpc"{
              aws_vpc := resources.Resources[i]
              for j := 0; j < len(aws_vpc.Instances); j++{
                  out = add_attr_vpc(aws_vpc.Instances[j].Attributes)

                  for l := 0; l < len(resources.Resources); l++ {
                    if resources.Resources[l].Type == "aws_subnet"{
                        aws_subnet := resources.Resources[l]
                        s := Subnet{
                          Name:aws_subnet.Name,
                          Image: "assets/img/aws_subnet.png",
                        }
                        for h := 0; h < len(aws_subnet.Instances); h++{
                          attr := aws_subnet.Instances[h].Attributes
                            s.Subnets = append(s.Subnets,add_attr(attr))
                        }
                        out.Subnets = append(out.Subnets, s)
                    }
                  }

              }

          }
          for l := 0; l < len(resources.Resources); l++ {
            if resources.Resources[l].Type == "aws_route"{
              aws_route := resources.Resources[l]
              w := RouteTable{
                Name: aws_route.Name,
              }
                for h := 0; h < len(aws_route.Instances); h++{
                    attr := aws_route.Instances[h].Attributes
                    w.Tables = append(w.Tables,add_attr_rt(attr))
                }
                out.RouteTables = append(out.RouteTables, w)
            }

          }

        }

    }


    var jsonData []byte
    jsonData, err = json.MarshalIndent(out, "", "  ")
    if isError(err) {
        return
    }
    id := uuid.New()
    output_path = output_path + id.String()
    ioutil.WriteFile(output_path, jsonData, 0777)
    fmt.Println("Started server...")

    start()
    err = os.Remove(output_path)
    if isError(err) {
        return
    }



    fmt.Println("Closed.")
}
