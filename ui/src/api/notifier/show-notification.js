import { ShowNotificationRequest } from 'notifier-api/gen/js/notifier_api_pb';

import { client } from './client';

export const showNotification = async ({ title, message }) => {
    const request = new ShowNotificationRequest();
    request.setTitle(title);
    request.setMessage(message);

    const response = await client.showNotification(request);

    return response.toObject();
}
