// file: src/routes/notebooks/[slug]/+page.server.ts

import { redirect, type ServerLoadEvent } from '@sveltejs/kit';

export type Notebook = {
	id: string;
	user_id: string;
	name: string;
};

export async function load(event: ServerLoadEvent) {
	const sessionId = event.cookies.get('sessionId') as string;

	if (!sessionId) {
		throw redirect(302, '/login');
	}

	return {
		sessionId,
		slug: event.params.slug
	};
}
