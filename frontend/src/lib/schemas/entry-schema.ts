import { z } from 'zod';

export const schema = z.object({
	content: z.string().min(1, 'Content cannot be empty').max(5000, 'Content is too long'),
	image_upload: z.string().optional()
});
