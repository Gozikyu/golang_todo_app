import { Suspense } from 'react';
import { Navigate, Outlet } from 'react-router-dom';

import { Spinner } from '@/components/Elements';
import { Tasks } from '@/features/tasks/routes/Tasks';
import AdminApp from '@/Admin/App';
import { MainLayout } from '@/components/Layout';

const App = () => {
  return (
    <MainLayout>
      <Suspense
        fallback={
          <div className="h-full w-full flex items-center justify-center">
            <Spinner size="xl" />
          </div>
        }
      >
        <Outlet />
      </Suspense>
    </MainLayout>
  );
};

export const protectedRoutes = [
  {
    element: <App />,
    children: [
      { path: '/tasks', element: <Tasks /> },
      { path: '/admin/*', element: <AdminApp /> },

      { path: '*', element: <Navigate to="." /> },
    ],
  },
];
