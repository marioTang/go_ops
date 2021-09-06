package utils

import (
	"cmdb/models"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func Ecscron() error{
	value := [...]string{"cn-qingdao",
		"cn-beijing",
		"cn-zhangjiakou",
		"cn-huhehaote",
		"cn-wulanchabu",
		"cn-hangzhou",
		"cn-shanghai",
		"cn-shenzhen",
		"cn-heyuan",
		"cn-guangzhou",
		"cn-chengdu",
		"cn-hongkong",
		"ap-northeast-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-southeast-3",
		"ap-southeast-5",
		"ap-south-1",
		"us-east-",
		"us-west-1",
		"eu-west-1",
		"me-east-1",
		"eu-central-1",
	}
	for _, regionid := range value {
		clientnumb, err := ecs.NewClientWithAccessKey(regionid, "LTAI4GBgFk1St2FgBdvbhgcb", "VdZkQhCk3IrdZllW1cnbXkpTnx1XcD")

		requestnumb := ecs.CreateDescribeInstancesRequest()
		requestnumb.Scheme = "https"
		//

		responsenumb, err := clientnumb.DescribeInstances(requestnumb)
		if err != nil {
			fmt.Print(err.Error())

		}
		datacount := (responsenumb.TotalCount)
		pagsize := datacount/30 + 1

		for i := 1; i <= pagsize; i++ {
			pagnumb := i

			client, err := ecs.NewClientWithAccessKey(regionid, "LTAI4GBgFk1St2FgBdvbhgcb", "VdZkQhCk3IrdZllW1cnbXkpTnx1XcD")

			request := ecs.CreateDescribeInstancesRequest()
			request.Scheme = "https"

			request.PageNumber = requests.NewInteger(pagnumb)
			request.PageSize = requests.NewInteger(30)

			response, err := client.DescribeInstances(request)
			if err != nil {
				fmt.Print(err.Error())

			}

			for _, v := range response.Instances.Instance {
				instanceid := models.SelectInstanceid(v.InstanceId)
				if v.InstanceId != instanceid.Instanceid {
					inneripaddress := v.VpcAttributes.PrivateIpAddress.IpAddress
					b, _ := json.Marshal(inneripaddress)
					var result = string(b)

					publicipaddress := v.PublicIpAddress.IpAddress
					c, _ := json.Marshal(publicipaddress)
					var result_c = string(c)

					data := models.Ecsinfo{1, v.InstanceId, v.InstanceName,
						v.InstanceType, v.Status, v.RegionId, v.CreationTime, v.HostName,
						v.CpuOptions.CoreCount, v.CpuOptions.ThreadsPerCore,
						v.InstanceChargeType, v.EipAddress.Bandwidth, v.EipAddress.IpAddress,
						result, result_c, v.InstanceNetworkType,
						v.LocalStorageAmount, "", v.Memory,
						v.OSName, v.OSType}
					go models.InsertEscinfo(data)

				}
			}
		}
	}

	return nil
}