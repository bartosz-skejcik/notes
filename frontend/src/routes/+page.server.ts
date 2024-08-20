import api from '$lib/api';
import { redirect, type ServerLoadEvent } from '@sveltejs/kit';

export type Notebook = {
	id: string;
	user_id: string;
	name: string;
	created_at: string;
};

type UserSession = {
	id: string;
	display_name: string;
	email: string;
};

const getSession = async (event: ServerLoadEvent): Promise<UserSession> => {
	const sessionId = event.cookies.get('sessionId');

	if (!sessionId) {
		return redirect(302, '/login');
	} else {
		const res = await api.get('/api/auth/me', {
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${sessionId}`
			}
		});

		if (res.status === 200) {
			return await res.data;
		} else {
			throw redirect(302, '/login');
		}
	}
};

const getNotebooks = async (sessionId: string): Promise<Notebook[]> => {
	const res = await api.get('/api/notebooks', {
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
	const user = await getSession(event);

	const sessionId = event.cookies.get('sessionId');

	if (!sessionId) {
		throw redirect(302, '/login');
	}

	const notebooks = await getNotebooks(sessionId);

	return {
		session: user,
		notebooks: notebooks
	};
}
