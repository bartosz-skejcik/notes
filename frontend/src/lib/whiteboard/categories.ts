import api from '$lib/api';

export interface Category {
	id: number;
	notebookId: number;
	name: string;
	color: string;
	noteCount: number;
}

export async function getCategories(notebookId: string, sessionId: string): Promise<Category[]> {
	const res = await api.get('/api/notebooks/' + notebookId + '/categories', {
		headers: {
			ContentType: 'application/json',
			Authorization: 'Bearer ' + sessionId
		}
	});

	if (res.status !== 200) {
		throw new Error('Failed to fetch categories');
	}

	return res.data;
}

export async function createCategory(
	notebookId: string,
	name: string,
	color: string,
	sessionId: string
): Promise<{ id: number } | null> {
	console.log('api create Category', notebookId, name, color, sessionId);
	const res = await api.post(
		'/api/notebooks/' + notebookId + '/categories',
		{
			name,
			color
		},
		{
			headers: {
				ContentType: 'application/json',
				Authorization: 'Bearer ' + sessionId
			}
		}
	);

	if (res.status !== 200) {
		throw new Error('Failed to create category');
	}

	return res.data;
}
