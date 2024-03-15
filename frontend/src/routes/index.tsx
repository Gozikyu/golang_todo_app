import { useRoutes } from 'react-router-dom';
import { protectedRoutes } from './protected';
import { publicRoutes } from './public';
import storage from '@/util/storage';
import Landing from '@/features/misc/routes/Landing';

export const AppRoutes = () => {
  const token = storage.getToken(); //TODO: ログイン中のユーザーを保持する機能を入れる

  const commonRoutes = [{ path: '/', element: <Landing /> }];
  const routes = token ? protectedRoutes : publicRoutes;

  const elements = useRoutes([...routes, ...commonRoutes]);

  return <>{elements}</>;
};
