const { host, protocol } = window.location;

export const GRPC_HOST = `${protocol}//${host}`;
