import { AuthRoutes } from '@/features/auth/routes';
import { Navigate } from 'react-router-dom';

export const publicRoutes = [
  { path: '*', element: <Navigate to="/auth/login" /> },
  {
    path: '/auth/*',
    element: <AuthRoutes />,
  },
];
