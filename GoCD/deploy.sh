#!/bin/bash -x
GITHASH=`git rev-parse --short HEAD`
S3KEY=GoCD-$GITHASH-STG_GoCD.zip

pwd
ls -la

aws s3 cp bin/GoCD_To_Deploy.zip s3://$S3BUCKET/$S3KEY
aws elasticbeanstalk create-application-version --application-name GoCD --version-label GoCD-$GITHASH-STG_GoCD --source-bundle S3Bucket="$S3BUCKET",S3Key="$S3KEY"
aws elasticbeanstalk update-environment --environment-name "gocd-env" --version-label GoCD-$GITHASH-STG_GoCD
