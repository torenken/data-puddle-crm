import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { DataPuddleCrmBucket } from './data-puddle-crm-bucket';
import { DataPuddleCrmEndpoint } from './endpoint/data-puddle-crm-endpoint';
import { TechnicalNotification } from './technical-notification';

export interface DataPuddleCrmStackProps extends StackProps {
  readonly emailAddresses: string[];
  readonly exportDataBucketName: string;
}

export class DataPuddleCrmStack extends Stack {
  constructor(scope: Construct, id: string, props: DataPuddleCrmStackProps) {
    super(scope, id, props);

    const notification = new TechnicalNotification(this, 'TechnicalNotification', {
      emailAddresses: props.emailAddresses,
    });

    const dataBucket = new DataPuddleCrmBucket(this, 'DataBucket', {
      bucketName: props.exportDataBucketName,
    });

    new DataPuddleCrmEndpoint(this, 'DataPuddleCrmEndpoint', {
      alarmNotification: notification,
      dataBucket,
    });
  }
}