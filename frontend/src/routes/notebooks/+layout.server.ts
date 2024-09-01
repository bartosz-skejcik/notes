import { error, redirect } from '@sveltejs/kit';

export async function load({ params, cookies }) {
  const sessionId = cookies.get('sessionId');

  if (!sessionId) {
    throw redirect(302, '/login');
  }

  if (!params.slug) {
    throw error(404, 'Not found');
  }

  return {
    sessionId,
    slug: params.slug,
  };
}
