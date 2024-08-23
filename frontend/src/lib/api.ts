import axios from 'axios';
import Cookies from 'js-cookie';

// Create an Axios instance with withCredentials set to true
const api = axios.create({
	baseURL: import.meta.env.VITE_API_URL
});

// Simple in-memory cache
const cache = new Map();

// Cache duration in milliseconds (e.g., 2 minutes)
const CACHE_DURATION = 30 * 1000;

api.interceptors.request.use(
	(config) => {
		// Get the cookie value
		const sessionIdCookie = Cookies.get('sessionId'); // Replace with your cookie name

		// If the cookie exists, attach it as a header
		if (sessionIdCookie) {
			config.headers['Authorization'] = `Bearer ${sessionIdCookie}`; // Adjust the header name and format as needed
			console.log(config.headers);
		}

		// Check cache for GET requests
		if (config.method && config.method.toLowerCase() === 'get' && config.url) {
			const cachedResponse = cache.get(config.url);
			if (cachedResponse) {
				const { data, timestamp } = cachedResponse;
				if (Date.now() - timestamp < CACHE_DURATION) {
					// Return cached data if it's still valid
					config.adapter = () =>
						Promise.resolve({
							data,
							status: 200,
							statusText: 'OK',
							headers: {},
							config,
							request: {}
						});
				}
			}
		}

		return config;
	},
	(error) => {
		return Promise.reject(error);
	}
);

api.interceptors.response.use(
	(response) => {
		// Cache successful GET responses
		if (
			response.config &&
			response.config.method &&
			response.config.method.toLowerCase() === 'get' &&
			response.config.url
		) {
			cache.set(response.config.url, {
				data: response.data,
				timestamp: Date.now()
			});
		}
		return response;
	},
	(error) => {
		return Promise.reject(error);
	}
);

export default api;
