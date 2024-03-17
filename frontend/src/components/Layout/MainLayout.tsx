import { ReactNode } from 'react';
import { AppBar, Toolbar, Typography, Button } from '@mui/material';
import { logOut } from '@/features/auth/api/logout';
import { useNavigate } from 'react-router-dom';

type MainLayoutProps = {
  children: ReactNode;
};

export const MainLayout = ({ children }: MainLayoutProps) => {
  const navigate = useNavigate();
  const handleLogout = () => {
    try {
      logOut();
      navigate('/');
      alert('ログアウトに成功しました');
    } catch (e) {
      alert('ログアウトに失敗しました');
      console.log(e);
    }
  };

  return (
    <div>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            ToDo App
          </Typography>
          <Button color="inherit" variant="outlined" onClick={handleLogout}>
            ログアウト
          </Button>
        </Toolbar>
      </AppBar>
      <main>{children}</main>
    </div>
  );
};
