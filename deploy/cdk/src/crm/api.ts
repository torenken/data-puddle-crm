import { Duration, RemovalPolicy } from 'aws-cdk-lib';
import { AccessLogFormat, EndpointType, LogGroupLogDestination, RestApi } from 'aws-cdk-lib/aws-apigateway';
import { SnsAction } from 'aws-cdk-lib/aws-cloudwatch-actions';
import { LogGroup, RetentionDays } from 'aws-cdk-lib/aws-logs';
import { ITopic } from 'aws-cdk-lib/aws-sns';
import { Construct } from 'constructs';

export interface ApiProps {
  readonly alarmNotification: ITopic;
}

export class Api extends RestApi {
  constructor(scope: Construct, id: string, props: ApiProps) {
    super(scope, id, {
      endpointConfiguration: {
        types: [EndpointType.REGIONAL],
      },
      deployOptions: {
        stageName: 'api',
        accessLogDestination: new LogGroupLogDestination(new LogGroup(scope, 'AccessLog', {
          retention: RetentionDays.THREE_MONTHS,
          removalPolicy: RemovalPolicy.DESTROY,
        })),
        accessLogFormat: AccessLogFormat.jsonWithStandardFields(),
        throttlingRateLimit: 3,
      },
    });

    const serverAlarm = this.metricServerError({ period: Duration.minutes(1) })
      .createAlarm(this, 'ApiMetrics5xAlarm', { alarmName: 'CrmApiMetrics5xAlarm', threshold: 1, evaluationPeriods: 2 });
    serverAlarm.addAlarmAction(new SnsAction(props.alarmNotification));

    const clientAlarm = this.metricClientError({ period: Duration.minutes(5) })
      .createAlarm(this, 'ApiMetrics4xAlarm', { alarmName: 'RiskApiMetrics4xAlarm', threshold: 3, evaluationPeriods: 1 });
    clientAlarm.addAlarmAction(new SnsAction(props.alarmNotification));

  }
}
