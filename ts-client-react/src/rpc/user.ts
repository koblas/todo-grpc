import { z } from "zod";

export const User = z.object({
  name: z.string(),
  email: z.string(),
});

export type UserType = z.infer<typeof User>;
