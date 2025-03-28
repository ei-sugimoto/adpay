import { z } from 'zod';

export const projectSchema = z.object({
    name: z.string().min(3).max(255),
});
