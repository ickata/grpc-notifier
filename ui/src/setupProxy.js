const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
	app.use(
		'/notifier.*',
		createProxyMiddleware({
			target: 'http://localhost:4333',
			changeOrigin: true,
			secure: false,
		})
	);
};
