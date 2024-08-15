import api from '$lib/api';
import { redirect, type ServerLoadEvent } from '@sveltejs/kit';

export async function load(event: ServerLoadEvent) {
	await api.post(
		'/api/auth/sign-out',
		{},
		{
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${event.cookies.get('sessionId')}`
			}
		}
	);

	event.cookies.delete('sessionId', { path: '/' });

	throw redirect(301, '/login');
}
