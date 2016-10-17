#!/usr/bin/python2.7
import time
import boto3
import os
import sys

def wait_deployment(client=None, **kwargs):
    if client is None:
        raise Exception('Client is mandatory')

    status = ''
    colour = ''
    older_version = client.describe_environments(EnvironmentNames=[kwargs['EnvironmentName'],])['Environments'][0]['VersionLabel']
    while not (colour == 'Green' and status == 'Ready'):
        try:
            print 'Fetching information from %s' % kwargs['EnvironmentName']
            response = client.describe_environment_health(**kwargs)
            if response['ResponseMetadata']['HTTPStatusCode'] == 200:
                colour = response['Color']
                status = response['Status']
                msg = 'Color: {}, Status: {}'
                print msg.format(colour, status)

        except Exception as e:
            print 'Error "{}" fetching data...'.format(e)
            raise e
        print 'Waiting 5 seconds...'
        time.sleep(5)
    current_version = client.describe_environments(EnvironmentNames=[kwargs['EnvironmentName'],])['Environments'][0]['VersionLabel']
    if current_version == older_version:
        print 'A rollback was peformed, running version is: %s' % current_version
        sys.exit(1)
    print 'Deployment was sucessful, running version is: %s' % current_version


def main():
    client = boto3.client('elasticbeanstalk')
    env_name = os.environ.get("AWS_DEPLOY_ENVIRONMENT")

    if env_name is None:
        print 'You must set "AWS_DEPLOY_ENVIRONMENT"!'
        sys.exit(1)

    kwargs = {
        'EnvironmentName': env_name,
        'AttributeNames': ['Status', 'Color'],
    }
    try:
        client = boto3.client('elasticbeanstalk')
        wait_deployment(client, **kwargs)
    except Exception as e:
        print e
        sys.exit(1)
    sys.exit(0)


if __name__ == "__main__":
    main()
