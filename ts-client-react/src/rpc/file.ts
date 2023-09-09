import { z } from "zod";

export const UploadUrlRequest = z.object({
  type: z.string(),
  contentType: z.string(),
});
export const UploadUrlResponse = z.object({
  url: z.string(),
  id: z.string(),
});
export type UploadUrlRequestT = z.infer<typeof UploadUrlRequest>;
export type UploadUrlResponseT = z.infer<typeof UploadUrlResponse>;

export const UploadFileRequest = z.object({
  url: z.string(),
  file: z.any(), // Technically a File
});
export const UploadFileResponse = z.unknown({});

export type UploadFileRequestT = z.infer<typeof UploadFileRequest>;
export type UploadFileResponseT = z.infer<typeof UploadFileResponse>;

export const FileEvent = z.object({
  object_id: z.string(),
  action: z.enum(["error", "create"]),
  topic: z.literal("file"),
  body: z.nullable(
    z.object({
      id: z.string(),
      error: z.optional(z.string()),
    }),
  ),
});
