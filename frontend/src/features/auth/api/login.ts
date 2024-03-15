import { axios } from '@/lib/axios';
import { LoginResponse, isLoginResponse } from '..';
import storage from '@/util/storage';

export type LoginVariables = {
  email: string;
  password: string;
};

const loginWithEmailAndPassword = async (
  data: LoginVariables
): Promise<LoginResponse> => {
  const res = await axios.post(`/login`, data);

  if (!isLoginResponse(res)) {
    throw new Error('responseがLoginResponseの型ではありません');
  }
  return res;
};

export const loginAndSetToken = async (data: LoginVariables) => {
  const response = await loginWithEmailAndPassword(data);
  const { token } = response;
  storage.setToken(token);
};
