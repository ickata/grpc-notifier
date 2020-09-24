import { GetNotificationsRequest } from 'notifier-api/gen/js/notifier_api_pb';

import { client } from './client';

export const getNotifications = async () => {
    const request = new GetNotificationsRequest();

    const response = await client.getNotifications(request);

    return response.toObject();
}