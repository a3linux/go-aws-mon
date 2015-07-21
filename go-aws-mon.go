package main

import (
	"strings"
    "encoding/json"
    "flag"
    "fmt"
    "log"
    "net/http"
	"io/ioutil"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
    "github.com/aws/aws-sdk-go/service/cloudwatch"
       )

func main() {

	isMemUtil := flag.Bool("mem-util", true, "Memory Utilization(percent)")
	isMemUsed := flag.Bool("mem-used", false, "Memory Used(bytes)")
	isMemAvail := flag.Bool("mem-avail", false, "Memory Available(bytes)")
	isSwapUtil := flag.Bool("swap-util", false, "Swap Utilization(percent)")
	isSwapUsed := flag.Bool("swap-used", false, "Swap Used(bytes)")
	isDiskSpaceUtil := flag.Bool("disk-space-util", true, "Disk Space Utilization(percent)")
	isDiskSpaceUsed := flag.Bool("disk-space-used", false, "Disk Space Used(bytes)")
	isDiskSpaceAvail := flag.Bool("disk-space-avail", false, "Disk Space Available(bytes)")
	isDiskInodeUtil := flag.Bool("disk-inode-util", false, "Disk Inode Utilization(percent)")

    ns := flag.String("namespace", "Linux/System", "CloudWatch metric namespace (required)(It is always EC2)")
	diskPaths := flag.String("disk-path", "/", "Disk Path")

    flag.Parse()

	paths := strings.Split(*diskPaths, ",")

	for k, val := range paths {
		fmt.Println(k)
		fmt.Println(val)
		fmt.Println(*isDiskSpaceAvail, *isDiskSpaceUsed, *isDiskSpaceUtil, *isDiskInodeUtil)
		diskspaceUtil, diskspaceUsed, diskspaceAvail, err := DiskSpace(val)
		if err != nil {
			log.Fatal("Can't get DiskSpace %s", err)
		}
		fmt.Println(diskspaceUtil, diskspaceAvail, diskspaceUsed)
	}

    memUtil, memUsed, memAvail, swapUtil, swapUsed, err := memoryUsage()

	if *isMemUtil {
		err = putMetric("MemoryUtilization", "Percent", memUtil, *ns)
		if err != nil {
    	    log.Fatal("Can't put memory usage metric: ", err)
    	}
	}

	if *isMemUsed {
		err = putMetric("MemoryUsed", "Bytes", memUsed, *ns)
		if err != nil {
			log.Fatal("Can't put memory used metric: ", err)
		}
	}
	if *isMemAvail {
		err = putMetric("MemoryAvail", "Bytes", memAvail, *ns)
		if err != nil {
			log.Fatal("Can't put memory available metric: ", err)
		}
	}
	if *isSwapUsed {
		err = putMetric("SwapUsed", "Bytes", swapUsed, *ns)
		if err != nil {
			log.Fatal("Can't put swap used metric: ", err)
		}
	}
	if *isSwapUtil {
		err = putMetric("SwapUtil", "Percent", swapUtil, *ns)
		if err != nil {
			log.Fatal("Can't put swap usage metric: ", err)
		}
	}

}

func putMetric(name, unit string, value float64, namespace string) error {

	region, instanceId, err := getMetadata()
	if err != nil {
		log.Fatal("Failed to get endpoint metadata: %s", err)
	}

	svc := cloudwatch.New(&aws.Config{Region: region})

	metric_input := &cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String(name),
				Unit:       aws.String(unit),
				Value:      aws.Double(value),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name: aws.String("InstanceId"),
						Value: aws.String(instanceId),
					},
				},
			},
		},
		Namespace: aws.String(namespace),
	}

	resp, err := svc.PutMetricData(metric_input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			return fmt.Errorf("[%s] %s", awsErr.Code, awsErr.Message)
		} else if err != nil {
			return err
		}
	}
	fmt.Println(awsutil.StringValue(resp))
	return nil
}

func getMetadata() (region string, instanceId string, err error) {
	resp, err := http.Get("http://169.254.169.254/latest/dynamic/instance-identity/document")
	if err != nil {
		return "", "", fmt.Errorf("can't reach metadata endpoint - %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("can't read metadata response body - %s", err)
	}

	var data map[string]string
	json.Unmarshal(body, &data)

	return data["region"], data["instanceId"], err
}
