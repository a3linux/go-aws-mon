package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)


func getAutoscalingGroup(instanceId string, region string) (*string, error) {
	session := session.New(&aws.Config{Region: &region})
	svc := autoscaling.New(session)

	params := &autoscaling.DescribeAutoScalingInstancesInput{
		InstanceIds: []*string{&instanceId},
		MaxRecords: aws.Int64(1),
	}

	resp, err := svc.DescribeAutoScalingInstances(params)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			return nil, fmt.Errorf("[%s] %s", awsErr.Code, awsErr.Message)
		} else if err != nil {
			return nil, err
		}
	}

	if len(resp.AutoScalingInstances) == 0 {
		return nil, errors.New("No autoscaling group found")
	}

	return resp.AutoScalingInstances[0].AutoScalingGroupName, nil
}