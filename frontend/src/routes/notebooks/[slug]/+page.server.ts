// file: src/routes/notebooks/[slug]/+page.server.ts

import api from '$lib/api';
import { redirect, type ServerLoadEvent } from '@sveltejs/kit';

type Notebook = {
	id: string;
	user_id: string;
	name: string;
};

const getNotebook = async (sessionId: string, slug: string): Promise<Notebook> => {
	const res = await api.get(`/api/notebooks/${slug}`, {
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${sessionId}`
		}
	});

	if (res.status === 200) {
		return res.data;
	} else {
		throw redirect(302, '/login');
	}
};

export async function load(event: ServerLoadEvent) {
	const sessionId = event.cookies.get('sessionId') as string;

	if (!sessionId) {
		throw redirect(302, '/login');
	}

	const slug = event.params.slug as string;

	const notebook = await getNotebook(sessionId, slug);

	return {
		sessionId,
		slug: event.params.slug,
		notebook
	};
}
