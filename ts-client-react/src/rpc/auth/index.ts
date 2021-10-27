export interface LoginSuccess {
  token: string;
}

export interface AuthService {
  register(params: { email: string; password: string; name: string; urlbase?: string }): Promise<LoginSuccess>;
  authenticate(email: string, password: string): Promise<LoginSuccess>;
}
