import { NotifierAPIPromiseClient } from 'notifier-api/gen/js/notifier_api_grpc_web_pb';

import { GRPC_HOST } from '../constants';

export const client = new NotifierAPIPromiseClient(GRPC_HOST);

if (window.__GRPCWEB_DEVTOOLS__) {
    window.__GRPCWEB_DEVTOOLS__([client]);
}
