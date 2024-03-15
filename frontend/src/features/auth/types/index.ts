export type LoginResponse = {
  token: string;
};

export function isLoginResponse(obj: any): obj is LoginResponse {
  return (
    typeof obj === 'object' &&
    obj !== null &&
    'token' in obj &&
    typeof obj.token === 'string'
  );
}
