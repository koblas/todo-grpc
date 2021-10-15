export type Json = string | number | boolean | null | JsonObject | JsonArray;

export interface JsonArray extends Array<Json> {}

// export type JsonObject = Record<string, Json>;
export interface JsonObject {
  [property: string]: Json;
}
