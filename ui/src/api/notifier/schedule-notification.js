import { ScheduleNotificationRequest } from 'notifier-api/gen/js/notifier_api_pb';

import { client } from './client';

export const scheduleNotification = async ({ title, message, datetime }) => {
    const request = new ScheduleNotificationRequest();
    request.setTitle(title);
    request.setMessage(message);
    request.setDatetime(Math.round(Number(datetime) / 1000));

    const response = await client.scheduleNotification(request);

    return response.toObject();
}
