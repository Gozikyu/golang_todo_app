const storagePrefix = 'todo_app_';

const storage = {
  getToken: (): string => {
    return JSON.parse(
      window.localStorage.getItem(`${storagePrefix}token`) as string
    );
  },
  getUserId: (): string => {
    return JSON.parse(
      window.localStorage.getItem(`${storagePrefix}userId`) as string
    );
  },
  setTokenAndUserId: (token: string, userId: string) => {
    window.localStorage.setItem(`${storagePrefix}token`, JSON.stringify(token));
    window.localStorage.setItem(
      `${storagePrefix}userId`,
      JSON.stringify(userId)
    );
  },
  clearTokenAndUserId: () => {
    window.localStorage.removeItem(`${storagePrefix}token`);
    window.localStorage.removeItem(`${storagePrefix}userId`);
  },
};

export default storage;
