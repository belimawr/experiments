#!/bin/bash -x
set -e
GITHASH=`git rev-parse --short HEAD`
S3KEY=$AWS_DEPLOY_APPLICATION-$GITHASH-$SUFFIX-$GO_PIPELINE_COUNTER.zip
VERSION_LABEL=$AWS_DEPLOY_APPLICATION-$GITHASH-$SUFFIX-$GO_PIPELINE_COUNTER

aws s3 cp bin/GoCD_To_Deploy.zip s3://$S3BUCKET/$S3KEY
aws elasticbeanstalk create-application-version --application-name $AWS_DEPLOY_APPLICATION --version-label $VERSION_LABEL --source-bundle S3Bucket="$S3BUCKET",S3Key="$S3KEY"
aws elasticbeanstalk update-environment --environment-name $AWS_DEPLOY_ENVIRONMENT --version-label $VERSION_LABEL
python ./wait.py
