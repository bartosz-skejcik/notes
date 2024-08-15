import axios from 'axios';
import Cookies from 'js-cookie';

// Create an Axios instance with withCredentials set to true
const api = axios.create({
	baseURL: import.meta.env.VITE_API_URL
});

api.interceptors.request.use(
	(config) => {
		// Get the cookie value
		const sessionIdCookie = Cookies.get('sessionId'); // Replace with your cookie name

		// If the cookie exists, attach it as a header
		if (sessionIdCookie) {
			config.headers['Authorization'] = `Bearer ${sessionIdCookie}`; // Adjust the header name and format as needed
			console.log(config.headers);
		}

		return config;
	},
	(error) => {
		return Promise.reject(error);
	}
);

export default api;
