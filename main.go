package main

// Sample expected Property file
// <?xml version="1.0" encoding="UTF-8"?>
// <Environment
//      xmlns="http://schemas.dmtf.org/ovf/environment/1"
//      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
//      xmlns:oe="http://schemas.dmtf.org/ovf/environment/1"
//      xmlns:ve="http://www.vmware.com/schema/ovfenv"
//      oe:id=""
//      ve:vCenterId="vm-621">
//    <PlatformSection>
//       <Kind>VMware ESXi</Kind>
//       <Version>6.5.0</Version>
//       <Vendor>VMware, Inc.</Vendor>
//       <Locale>en</Locale>
//    </PlatformSection>
//    <PropertySection>
//          <Property oe:key="DNS" oe:value="10.193.190.2"/>
//          <Property oe:key="admin_password" oe:value="admin"/>
//          <Property oe:key="gateway" oe:value="172.16.20.1"/>
//          <Property oe:key="ip0" oe:value="172.16.20.5"/>
//          <Property oe:key="netmask0" oe:value="255.255.255.0"/>
//          <Property oe:key="ntp_servers" oe:value="10.193.190.2"/>
//    </PropertySection>
//    <ve:EthernetAdapterSection>
//    </ve:EthernetAdapterSection>
// </Environment>

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Environment struct {
	PropertySection PropertySection
}

type PropertySection struct {
	Properties []Property `xml:"Property"`
}

type Property struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var output []byte

	for {
		input, e := reader.ReadByte()
		if e != nil && e == io.EOF {
			break
		}
		output = append(output, input)
	}

	var environ Environment
	err := xml.Unmarshal(output, &environ)
	if err != nil {
		fmt.Println(err)
		return
	}
	m := make(map[string]string)
	for _, p := range environ.PropertySection.Properties {
		m[p.Key] = p.Value
	}
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s", m["DNS"], m["admin_password"], m["gateway"], m["ip0"], m["netmask0"], m["ntp_servers"])
	fmt.Println()
}
