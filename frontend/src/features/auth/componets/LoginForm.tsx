import { useState } from 'react';
import { loginAndSetToken } from '../api/login';
import { useNavigate } from 'react-router-dom';
import { TextField, Button, Container, Grid, Typography } from '@mui/material';

const LoginForm = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      await loginAndSetToken({ email, password });
      alert('ログインに成功しました');
      navigate('/tasks');
    } catch (e) {
      alert('ログインに失敗しました');
    }
  };

  return (
    <Container maxWidth="sm">
      <Grid container spacing={2} justifyContent="center">
        <Grid item xs={12}>
          <Typography variant="h5" align="center" gutterBottom>
            ログイン
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            type="email"
            label="メールアドレス"
            variant="outlined"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            type="password"
            label="パスワード"
            variant="outlined"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </Grid>
        <Grid item xs={4}>
          <Button
            fullWidth
            variant="contained"
            color="primary"
            onClick={handleLogin}
          >
            ログイン
          </Button>
        </Grid>
      </Grid>
    </Container>
  );
};

export default LoginForm;
