import {
	getStickyNotes,
	getStickyNotesForUserCategory,
	type StickyNote
} from '$lib/whiteboard/sticky_notes.js';
import { error, redirect, type ServerLoadEvent } from '@sveltejs/kit';

export async function load({ url, params, cookies }: ServerLoadEvent) {
	const view = url.searchParams.get('view');

	const sessionId = cookies.get('sessionId');

	if (!sessionId) {
		throw redirect(302, '/login');
	}

	if (!params.slug) {
		throw error(404, 'Not found');
	}

	let sticky_notes: StickyNote[] = [];

	if (view && view === 'all') {
		sticky_notes = await getStickyNotes(params.slug, sessionId);
	} else if (view && view === 'recent') {
		const notes = await getStickyNotes(params.slug, sessionId);
		sticky_notes = notes.sort(
			(a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
		);
	} else {
		if (view) {
			sticky_notes = await getStickyNotesForUserCategory(params.slug, view, sessionId);
		}
	}

	return {
		view,
		sticky_notes
	};
}
