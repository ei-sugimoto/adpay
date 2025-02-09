import { z } from 'zod';

export const loginSchema = z.object({
  name: z
    .string()
    .min(3)
    .max(255)
    .regex(/^[a-zA-Z0-9]+$/),
  password: z
    .string()
    .min(8)
    .max(255)
    .regex(/^[a-zA-Z0-9]+$/),
});
