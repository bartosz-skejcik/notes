import type { RequestHandler } from './$types';
import { env } from '$env/dynamic/private';

import { StreamData, streamText } from 'ai';
import { createOpenAI as createGroq } from '@ai-sdk/openai';

const groq = createGroq({
	baseURL: 'https://api.groq.com/openai/v1',
	apiKey: env.GROQ_API_KEY
});
export const POST = (async ({ request }) => {
	const { messages } = await request.json();

	const data = new StreamData();
	data.append({ test: 'value' });

	const result = await streamText({
		model: groq('llama-3.1-70b-versatile'),
		onFinish() {
			data.close();
		},
		system:
			'You are an AI within a journaling app. Your job is to help the user reflect on their thoughts in a thoughtful and kind manner. The user can never directly address you or directly respond to you. Try not to repeat what the user said, instead try to seed new ideas, encourage or debate. Keep your responses concise, but meaningful.',
		messages
	});

	return result.toDataStreamResponse();
}) satisfies RequestHandler;
