package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform/terraform"
)

type StateVersions struct {
	Versions map[string]*terraform.State
}

func GetState(st, versionId string) (state *terraform.State, err error) {
	log.Infof("Getting state for %s/%s", st, versionId)
	if _, ok := stateVersions[st]; !ok {
		// Init
		stateVersions[st] = &StateVersions{}
		stateVersions[st].Versions = make(map[string]*terraform.State)
	}

	if s, ok := stateVersions[st].Versions[versionId]; ok {
		// Return cached version
		log.Infof("Getting %s/%s from cache", st, versionId)
		return s, nil
	}

	// Retrieve from S3
	log.Infof("Retrieving %s/%s from S3", st, versionId)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(st),
	}
	if versionId != "" {
		input.VersionId = &versionId
	}
	result, err := svc.GetObjectWithContext(context.Background(), input)
	if err != nil {
		log.Errorf("Error retrieving %s/%s from S3: %v", st, versionId, err)
		errObj := make(map[string]string)
		errObj["error"] = fmt.Sprintf("State file not found: %v", st)
		errObj["details"] = fmt.Sprintf("%v", err)
		j, _ := json.Marshal(errObj)
		return state, fmt.Errorf("%s", string(j))
	}
	defer result.Body.Close()

	content, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Errorf("Error reading %s/%s from S3: %v", st, versionId, err)
		errObj := make(map[string]string)
		errObj["error"] = fmt.Sprintf("Failed to read S3 response: %v", st)
		errObj["details"] = fmt.Sprintf("%v", err)
		j, _ := json.Marshal(errObj)
		return state, fmt.Errorf("%s", string(j))
	}

	json.Unmarshal(content, &state)

	if versionId != "" {
		// Store in cache
		log.Infof("Adding %s/%s to cache", st, versionId)
		stateVersions[st].Versions[versionId] = state
	}

	return
}