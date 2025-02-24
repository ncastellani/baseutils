package baseutils

import "os"

// - region translations

// application region to AWS region map
var RegionToAWSRegion map[string]string = map[string]string{
	"saopaulo": "sa-east-1",
	"montreal": "ca-central-1",
	"virginia": "us-east-1",
}

// AWS region to application region map
var AWSRegionToRegion map[string]string = map[string]string{
	"sa-east-1":    "saopaulo",
	"ca-central-1": "montreal",
	"us-east-1":    "virginia",
}

// Fly.io region into application region name
var FlyioRegionToAPPRegion map[string]string = map[string]string{
	"gru": "saopaulo",
	"yul": "montreal",
	"iad": "virginia",
}

// - Fly.io utils

// determine if the current running instance is a Fly.io instance
func IsFlyioInstance() (isFly bool, region string) {
	if os.Getenv("FLY_MACHINE_ID") != "" {
		return true, FlyioRegionToAPPRegion[os.Getenv("FLY_REGION")]
	}

	return
}
