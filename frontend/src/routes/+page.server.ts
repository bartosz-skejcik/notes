import api from '$lib/api';
import { redirect, type ServerLoadEvent } from '@sveltejs/kit';

type Data = {
	notebooks: {
		id: string;
		icon: string;
		name: string;
	}[];
};

const d: Data = {
	notebooks: [
		{
			id: '1',
			icon: '📓',
			name: 'Personal'
		},
		{
			id: '2',
			icon: '📣',
			name: 'Social media'
		}
	]
};

type UserSession = {
	id: string;
	display_name: string;
	email: string;
};

const getSession = async (event: ServerLoadEvent): Promise<UserSession | undefined> => {
	const sessionId = event.cookies.get('sessionId');

	if (!sessionId) {
		throw redirect(302, '/login');
	} else {
		const res = await api.get('/api/auth/me', {
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
	}
};

export async function load(event: ServerLoadEvent) {
	const session = await getSession(event);

	return {
		session,
		notebooks: d.notebooks
	};
}
