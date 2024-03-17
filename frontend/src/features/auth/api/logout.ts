import storage from '@/util/storage';

export const logOut = () => {
  storage.clearTokenAndUserId();
};
