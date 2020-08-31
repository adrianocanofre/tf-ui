package main


type Resources struct {
    Resources []Resource
}

type Resource struct {
    Type  string
    Name string
    Instances []Intance
}

type Intance struct {
  Index int
  Attributes Attribute
}

type Attribute struct {
  Vpc_id string
  Cidr_block string
  Name string
  Id string
  Arn string
  Tags Tag
  Gateway_id string
  Route_table_id string
  Nat_gateway_id string
  Transit_gateway_id string
  Default_security_group_id string
  Network_interface string
  Private_dns string
  Private_ip string
  Public_dns string
  Public_ip string
}

type Tag struct {
  Name string
}

type Vpc struct {
  Image string `json:"image"`
  Tags string `json:"vpc_name"`
  Id string `json:"vpc_id"`
  Cidr_block string `json:"cidr_block"`
  Default_security_group_id string `json:"security_group_id"`
  Subnets []Subnet `json:"subnets"`
  RouteTables []RouteTable `json:"route_table"`

}

type Subnet struct {
  Image string `json:"image"`
  Name string `json:"name"`
  Subnets []SubnetData `json:"subnets"`
}

type SubnetData struct {
  Arn string `json:"arn"`
  Cidr_block string `json:"cidr_block"`
  Id string `json:"id"`
  Tags string `json:"subnet_name"`
}

type RouteTable struct {
  Name string `json:"name"`
  Tables []RouteTableData `json:"route_table"`

}

type RouteTableData struct {
  Id string `json:"id"`
  RouteTableId string `json:"route_table_id"`
  NatGatewayId string `json:"nat_gateway_id"`
  GatewayId string `json:"gateway_id"`
  TransitGatewayId string `json:"transit_gateway_id"`
}
