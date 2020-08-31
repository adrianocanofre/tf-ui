package main

import (
  "log"
)

func add_attr(attr Attribute) SubnetData{
  w := SubnetData{
    Arn: attr.Arn,
    Id: attr.Id,
    Cidr_block: attr.Cidr_block,
    Tags: attr.Tags.Name,
  }


  return w
}


func add_attr_rt(attr Attribute) RouteTableData{
  w := RouteTableData{
    Id: attr.Id,
    RouteTableId: attr.Route_table_id,
    NatGatewayId: attr.Nat_gateway_id,
    GatewayId: attr.Gateway_id,
    TransitGatewayId: attr.Transit_gateway_id,
  }


  return w
}


func add_attr_vpc(attr Attribute) Vpc{
  w := Vpc{
    Image: path_assets_img + "aws_vpc.png",
    Id: attr.Id,
    Cidr_block: attr.Cidr_block,
    Tags: attr.Tags.Name,
    Default_security_group_id: attr.Default_security_group_id,
  }


  return w
}

func isError(err error) bool {
    if err != nil {
        log.Println(err)
    }
    return (err != nil)
}


// func openFile(file string) (jsonFile *os.File) {
//   jsonFile, err := os.Open("terraform.tfstate")
//
//   if err != nil {
//       fmt.Println(err)
//   }
//   fmt.Println("Successfully Opened users.json")
//   defer jsonFile.Close()
//   return jsonFile
// }
