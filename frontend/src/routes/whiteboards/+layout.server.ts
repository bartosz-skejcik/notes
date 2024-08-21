import { getCategories } from '$lib/whiteboard/categories.js';
import { error, redirect } from '@sveltejs/kit';

export async function load({ cookies, params }) {
	const sessionId = cookies.get('sessionId');

	if (!sessionId) {
		throw redirect(302, '/login');
	}

	if (!params.slug) {
		throw error(404, 'Not found');
	}

	console.log(params.slug);

	const categories = await getCategories(params.slug, sessionId);

	console.log(categories);

	return {
		sessionId,
		slug: params.slug,
		categories
	};
}
