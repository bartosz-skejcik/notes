import { z } from 'zod';

// Define outside the load function so the adapter can be cached
export const schema = z.object({
	firstName: z.string().min(3),
	lastName: z.string().min(3),
	email: z.string().email(),
	password: z.string().min(6)
});
