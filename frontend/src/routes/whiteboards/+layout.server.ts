import { getCategories } from '$lib/whiteboard/categories.js';
import { error, redirect, type ServerLoadEvent } from '@sveltejs/kit';

export async function load({ cookies, params, url }: ServerLoadEvent) {
	const sessionId = cookies.get('sessionId');

	if (!sessionId) {
		throw redirect(302, '/login');
	}

	if (!params.slug) {
		throw error(404, 'Not found');
	}

	// get the query param named whiteboard

	const view = url.searchParams.get('view');

	const categories = await getCategories(params.slug, sessionId);

	return {
		sessionId,
		slug: params.slug,
		categories,
		view: view ?? 'recent'
	};
}
