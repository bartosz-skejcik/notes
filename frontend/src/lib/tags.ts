import api from './api';

export interface Tag {
	id: number;
	label: string;
	value: string;
	color: string;
}

export async function fetchTags(sessionId: string): Promise<Tag[] | null> {
	const res = await api.get('/api/tags', {
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${sessionId}`
		}
	});
	if (res.status === 200) {
		return res.data;
	} else {
		return null;
	}
}
