import { z } from 'zod';

// Define outside the load function so the adapter can be cached
export const schema = z.object({
	email: z.string().email(),
	password: z.string().min(6)
});
