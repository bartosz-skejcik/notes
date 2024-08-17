import api from '$lib/api';
import { redirect, type Actions, type RequestEvent, type ServerLoadEvent } from '@sveltejs/kit';

import { fail, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { schema } from '$lib/schemas/new-notebook-schema';

type UserSession = {
	id: string;
	display_name: string;
	email: string;
};

const getSession = async (event: ServerLoadEvent): Promise<UserSession | undefined> => {
	const sessionId = event.cookies.get('sessionId')!;

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

	const form = await superValidate(zod(schema));

	return {
		form,
		session
	};
}

export const actions: Actions = {
	default: async (event: RequestEvent) => {
		const sessId = event.cookies.get('sessionId');
		const form = await superValidate(event.request, zod(schema));

		if (!form.valid) {
			return fail(400, { form });
		}

		try {
			const res = await api.post(
				'/api/notebooks',
				{
					name: form.data.name
				},
				{
					headers: {
						'Content-Type': 'application/json',
						Authorization: `Bearer ${sessId}`
					}
				}
			);

			if (res.status === 200) {
				return redirect(302, '/');
			} else {
				throw new Error(`Unexpected response status: ${res.status}`);
			}
		} catch (error) {
			// @ts-expect-error asd
			if (error.response) {
				// @ts-expect-error asd
				console.error('Request failed with status code:', error.response.status);
				// @ts-expect-error asd
				console.error('Response body:', error.response.data);

				// @ts-expect-error asd
				throw redirect(301, `/new-notebook?error=${error.response.data.message}`);
			} else {
				// @ts-expect-error asd
				console.error('Request error:', error.message);
				throw redirect(301, '/');
			}
		}
	}
};
