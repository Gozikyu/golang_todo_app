import Axios, {
  AxiosError,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from 'axios';

import { API_URL, UNAUTHENTICATED_API_URL } from '@/config';
import storage from '@/util/storage';

export const axios = Axios.create({
  baseURL: API_URL,
});

export const axiosForUnauthenticated = Axios.create({
  baseURL: UNAUTHENTICATED_API_URL,
});

const authRequestInterceptor = (config: InternalAxiosRequestConfig) => {
  const token = storage.getToken();
  if (!config.headers) {
    return config;
  }

  if (token) {
    config.headers.authorization = `Bearer ${token}`;
  }
  config.headers.Accept = 'application/json';
  return config;
};

const responseInterceptor = (response: AxiosResponse) => {
  return response.data;
};

const errorInterceptor = (error: AxiosError) => {
  return Promise.reject(error);
};

axios.interceptors.request.use(authRequestInterceptor);
axios.interceptors.response.use(responseInterceptor, errorInterceptor);
axiosForUnauthenticated.interceptors.response.use(
  responseInterceptor,
  errorInterceptor
);
