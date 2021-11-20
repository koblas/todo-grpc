export function randomString(length: number): string {
  const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";

  return new Array(length).map(() => possible.charAt(Math.floor(Math.random() * possible.length))).join("");
}
