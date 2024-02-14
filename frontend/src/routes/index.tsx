import { useRoutes } from 'react-router-dom';
import { protectedRoutes } from './protected';
import { Landing } from '@/features/misc';

export const AppRoutes = () => {
  const commonRoutes = [{ path: '/', element: <Landing /> }];

  const routes = protectedRoutes; //TODO: 認証済み、未認証でrouteを分岐させるようにする
  const elements = useRoutes([...routes, ...commonRoutes]);

  return <>{elements}</>;
};
