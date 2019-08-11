//AWS Golang code to list all instance under a specific region based on aws config


package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)


func main() {
    // Load session from shared config
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    svc := ec2.New(sess)


    filters := []*ec2.Filter{
    		&ec2.Filter{
    			Name: aws.String("tag:Name"),
    			Values: []*string{
    				aws.String("*"),
    			},
    		},
    	}


      input := ec2.DescribeInstancesInput{Filters: filters}
      result, err := svc.DescribeInstances(&input)
      if err != nil {
		panic(err.Error())
	}


  for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {

         fmt.Printf("%s", *instance.InstanceId)
         fmt.Println()
		}
	}


}
