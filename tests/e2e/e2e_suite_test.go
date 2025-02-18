package e2e

import (
	"flag"
	"log"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Common vars obtained from flags passed in ginkgo.
var cloud, namespace, s3Bucket, s3BucketFilePath, credSecretRef, instanceName, region, provider string
var timeoutMultiplier time.Duration

func init() {
	flag.StringVar(&cloud, "cloud", "", "Cloud Credentials file path location")
	flag.StringVar(&s3BucketFilePath, "s3_bucket", "myBucket", "AWS S3 data file path location")
	flag.StringVar(&namespace, "velero_namespace", "oadp-operator", "Velero Namespace")
	flag.StringVar(&region, "region", "us-east-1", "BSL region")
	flag.StringVar(&provider, "provider", "aws", "BSL provider")
	flag.StringVar(&credSecretRef, "creds_secret_ref", "cloud-credentials", "Credential secret ref for backup storage location")
	flag.StringVar(&instanceName, "velero_instance_name", "example-velero", "Velero Instance Name")
	timeoutMultiplierInput := flag.Int64("timeout_multiplier", 1, "Customize timeout multiplier from default (1)")
	timeoutMultiplier = 1
	if timeoutMultiplierInput != nil && *timeoutMultiplierInput >= 1 {
		timeoutMultiplier = time.Duration(*timeoutMultiplierInput)
	}
}

func TestOADPE2E(t *testing.T) {
	flag.Parse()
	s3Buffer, err := getJsonData(s3BucketFilePath)
	if err != nil {
		t.Fatalf("Error getting bucket json file: %v", err)
	}
	s3Data, err := decodeJson(s3Buffer) // Might need to change this later on to create s3 for each tests
	if err != nil {
		t.Fatalf("Error decoding json file: %v", err)
	}
	s3Bucket = s3Data["velero-bucket-name"].(string)
	log.Println("Using velero prefix: " + veleroPrefix)
	RegisterFailHandler(Fail)
	RunSpecs(t, "OADP E2E Suite")
}
