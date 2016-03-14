package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"io/ioutil"
	"log"
	"net/http"
)

func getDimensions(metadata map[string]string) (ret []*cloudwatch.Dimension) {

	var _ret []*cloudwatch.Dimension

	instanceIdName := "InstanceId"
	instanceIdValue, ok := metadata["instanceId"]
	if ok {
		dim := cloudwatch.Dimension{
			Name:  aws.String(instanceIdName),
			Value: aws.String(instanceIdValue),
		}
		_ret = append(_ret, &dim)
	}

	imageIdName := "ImageId"
	imageIdValue, ok := metadata["imageId"]
	if ok {
		dim := cloudwatch.Dimension{
			Name:  aws.String(imageIdName),
			Value: aws.String(imageIdValue),
		}
		_ret = append(_ret, &dim)
	}

	instanceTypeName := "InstanceType"
	instanceTypeValue, ok := metadata["instanceType"]
	if ok {
		dim := cloudwatch.Dimension{
			Name:  aws.String(instanceTypeName),
			Value: aws.String(instanceTypeValue),
		}
		_ret = append(_ret, &dim)
	}

	fileSystemName := "FileSystem"
	fileSystemValue, ok := metadata["fileSystem"]
	if ok {
		dim := cloudwatch.Dimension{
			Name:  aws.String(fileSystemName),
			Value: aws.String(fileSystemValue),
		}
		_ret = append(_ret, &dim)
	}

	return _ret
}

func addMetric(name, unit string, value float64, dimensions []*cloudwatch.Dimension, metricData []*cloudwatch.MetricDatum) (ret []*cloudwatch.MetricDatum, err error) {
	_metric := cloudwatch.MetricDatum{
		MetricName: aws.String(name),
		Unit:       aws.String(unit),
		Value:      aws.Float64(value),
		Dimensions: dimensions,
	}
	metricData = append(metricData, &_metric)
	return metricData, nil
}

func putMetric(metricdata []*cloudwatch.MetricDatum, namespace, region string) error {

	session := session.New(&aws.Config{Region: &region})
	svc := cloudwatch.New(session)

	metric_input := &cloudwatch.PutMetricDataInput{
		MetricData: metricdata,
		Namespace:  aws.String(namespace),
	}

	resp, err := svc.PutMetricData(metric_input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			return fmt.Errorf("[%s] %s", awsErr.Code, awsErr.Message)
		} else if err != nil {
			return err
		}
	}
	log.Println(awsutil.StringValue(resp))
	return nil
}

/*
  Metadata struct:
  {
    "devpayProductCodes" : null,
	"privateIp" : "10.0.5.89",
	"availabilityZone" : "us-west-1a",
	"version" : "2010-08-31",
	"region" : "us-west-1",
	"instanceId" : "i-e0iag2b",
	"billingProducts" : null,
	"accountId" : "208372078340",
	"instanceType" : "m3.xlarge",
	"imageId" : "ami-43f91b07",
	"kernelId" : null,
    "ramdiskId" : null,
    "pendingTime" : "2015-06-30T08:28:48Z",
    "architecture" : "x86_64"
  }
*/
func getInstanceMetadata() (metadata map[string]string, err error) {
	var data map[string]string
	resp, err := http.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	if err != nil {
		return data, fmt.Errorf("can't reach metadata endpoint - %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, fmt.Errorf("can't read metadata response body - %s", err)
	}

	json.Unmarshal(body, &data)

	return data, err
}
