import api from '$lib/api';
import {
	fail,
	redirect,
	type Actions,
	type RequestEvent,
	type ServerLoadEvent
} from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { schema } from '$lib/schemas/register-schema';

export async function load(event: ServerLoadEvent) {
	const sessionId = event.cookies.get('sessionId');

	if (sessionId) {
		throw redirect(301, '/');
	}

	// get the path query parameter
	const error = event.url.searchParams.get('error');

	const form = await superValidate(zod(schema));

	return {
		error,
		form
	};
}

export const actions: Actions = {
	default: async (event: RequestEvent) => {
		const form = await superValidate(event.request, zod(schema));

		if (!form.valid) {
			return fail(400, { form });
		}

		const displayName = `${form.data.firstName} ${form.data.lastName}`;

		const data = {
			display_name: displayName,
			email: form.data.email,
			password: form.data.password
		};

		try {
			const res = await api.post('/api/auth/sign-up', data, {
				headers: { 'Content-Type': 'application/json' }
			});

			if (res.status === 200) {
				const header = res.headers['authorization'];
				const sessionId = header?.split(' ')[1];

				event.cookies.set('sessionId', sessionId, { path: '/' });
				throw redirect(301, '/');
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
				throw redirect(301, `/register?error=${error.response.data.message}`);
			} else {
				// @ts-expect-error asd
				console.error('Request error:', error.message);
				throw redirect(301, '/');
			}
		}
	}
};
