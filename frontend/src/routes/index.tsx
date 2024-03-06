import { useRoutes } from 'react-router-dom';
import { protectedRoutes } from './protected';
import { Landing } from '@/features/misc';
import { publicRoutes } from './public';

export const AppRoutes = () => {
  const user = 'user'; //TODO: ログイン中のユーザーを保持する機能を入れる
  const commonRoutes = [{ path: '/', element: <Landing /> }];

  const routes = user ? protectedRoutes : publicRoutes; //TODO: 認証済み、未認証でrouteを分岐させるようにする
  const elements = useRoutes([...routes, ...commonRoutes]);

  return <>{elements}</>;
};
