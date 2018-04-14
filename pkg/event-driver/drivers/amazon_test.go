///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package drivers

import (
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func TestCloudWatchEventRuleEnable(t *testing.T) {
	sess := getSession()

	svc := cloudwatchevents.New(sess)

	output, err := svc.EnableRule(&cloudwatchevents.EnableRuleInput{
		Name: aws.String("Dispatch-Test"),
	})

	if err != nil {
		t.Logf("the error was %s", err.Error())
	}

	t.Logf("enable rule output: %s", output.String())
}

func TestCloudWatchEventRuleDisable(t *testing.T) {
	sess := getSession()

	svc := cloudwatchevents.New(sess)

	output, err := svc.DisableRule(&cloudwatchevents.DisableRuleInput{
		Name: aws.String("Dispatch-Test"),
	})

	if err != nil {
		t.Logf("the error was %s", err.Error())
	}

	t.Logf("disable rule output: %s", output.String())
}

func TestCloudWatchEventPutRule(t *testing.T) {
	sess := getSession()

	svc := cloudwatchevents.New(sess)

	output, err := svc.PutRule(&cloudwatchevents.PutRuleInput{
		State:              aws.String(cloudwatchevents.RuleStateDisabled),
		Name:               aws.String("Test-PutRule"),
		Description:        aws.String("Test of GO API for creating rules"),
		ScheduleExpression: aws.String("rate(5 minutes)"),
	})

	if err != nil {
		t.Logf("the error was %s", err.Error())
	}

	t.Logf("put rule output: %s", output.String())

	assert.Fail(t, "Just capturing logs")
}

func TestCloudWatchEventDeleteRule(t *testing.T) {
	sess := getSession()

	svc := cloudwatchevents.New(sess)

	output, err := svc.DeleteRule(&cloudwatchevents.DeleteRuleInput{
		Name: aws.String("Test-PutRule"),
	})

	if err != nil {
		t.Logf("the error was %s", err.Error())
	}

	t.Logf("delete rule output: %s", output.String())

	assert.Fail(t, "Just capturing logs")
}

func TestCloudWatchEventPutTarget(t *testing.T) {
	sess := getSession()

	svc := cloudwatchevents.New(sess)

	output, err := svc.PutTargets(&cloudwatchevents.PutTargetsInput{
		Rule: aws.String("Test-PutRule"),
		Targets: []*cloudwatchevents.Target{&cloudwatchevents.Target{
			Id:  aws.String("dispatch_test_sqs"),
			Arn: aws.String("arn:aws:sqs:us-west-2:367199020685:dispatch_test"),
		}},
	})

	if err != nil {
		t.Logf("the error was %s", err.Error())
	}

	t.Logf("put target output: %s", output.String())

	assert.Fail(t, "Just capturing logs")
}

func TestCloudWatchEventDeleteTarget(t *testing.T) {
	sess := getSession()

	svc := cloudwatchevents.New(sess)

	output, err := svc.RemoveTargets(&cloudwatchevents.RemoveTargetsInput{
		Ids:  []*string{aws.String("dispatch_test_sqs")},
		Rule: aws.String("Test-PutRule"),
	})

	if err != nil {
		t.Logf("the error was %s", err.Error())
	}

	t.Logf("remove target output: %s", output.String())

	assert.Fail(t, "Just capturing logs")
}
func TestAmazon(t *testing.T) {

	sess := getSession()

	metricName := "numberOfMessagesPublished"

	cloudwatchSvc := cloudwatch.New(sess)
	output, err := cloudwatchSvc.ListMetrics(&cloudwatch.ListMetricsInput{})

	metrics := output.Metrics
	for _, metric := range metrics {
		t.Logf("metric name: %s, metric namespace: %s", *metric.MetricName, *metric.Namespace)
	}
	now := time.Now()
	then := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	out, err := cloudwatchSvc.GetMetricData(&cloudwatch.GetMetricDataInput{
		EndTime:       &now,
		MaxDatapoints: &[]int64{100}[0],
		MetricDataQueries: []*cloudwatch.MetricDataQuery{&cloudwatch.MetricDataQuery{
			Id: &metricName,
			MetricStat: &cloudwatch.MetricStat{
				Metric: &cloudwatch.Metric{},
				Period: &[]int64{1}[0],
				Stat:   &metricName,
			},
		}},
		StartTime: &then,
	})

	t.Logf("The errors are %+v", err)
	assert.Nil(t, err, "Received error getting metrics")

	assert.Empty(t, out.MetricDataResults, "had results")
	fmt.Println("the endpoint is not nil")
}

func getSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String("us-west-2"),
			Credentials: credentials.NewSharedCredentials("/Users/ntenczar/.aws/credentials", "ntenczar"),
		},
	}))
}
