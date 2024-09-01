// file: src/stores/notebook.svelte.ts

import api from '$lib/api';
import { redirect } from '@sveltejs/kit';
import type { Notebook } from '../routes/+page.server';

let notebooks: Notebook[] = $state([]);
let notebook: Notebook | undefined = $state();

export function useNotebooksStore() {
	async function getNotebooks(sessionId: string): Promise<Notebook[]> {
		const res = await api.get('/api/notebooks', {
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${sessionId}`
			}
		});

		if (res.status === 200) {
			notebooks = res.data;
			return notebooks;
		} else {
			throw redirect(302, '/login');
		}
	}

	async function getNotebook(sessionId: string, slug: string): Promise<Notebook> {
		const res = await api.get(`/api/notebooks/${slug}`, {
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${sessionId}`
			}
		});

		if (res.status === 200) {
			const fetchedNotebook = res.data;

			notebook = fetchedNotebook;

			return fetchedNotebook;
		} else {
			throw redirect(302, '/login');
		}
	}

	return {
		get notebooks() {
			return notebooks;
		},
		get notebook() {
			return notebook;
		},
		getNotebooks,
		getNotebook
	};
}
