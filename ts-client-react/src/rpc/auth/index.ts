export interface LoginSuccess {
  token: string;
}

export interface AuthService {
  login(username: string, password: string): Promise<LoginSuccess>;
}
