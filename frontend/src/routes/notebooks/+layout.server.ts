import { fetchEntries } from '$lib/entry.js';
import { error, redirect } from '@sveltejs/kit';

export async function load({ params, cookies }) {
	const sessionId = cookies.get('sessionId');

	if (!sessionId) {
		throw redirect(302, '/login');
	}

	if (!params.slug) {
		throw error(404, 'Not found');
	}

	const entries = await fetchEntries({ sessionId: sessionId, slug: params.slug });

	console.log(entries);

	return {
		entries
	};
}
