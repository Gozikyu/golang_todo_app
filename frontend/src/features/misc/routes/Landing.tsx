import { Button, Container, Typography, Grid } from '@mui/material';
import { Link } from 'react-router-dom';

const Landing = () => {
  return (
    <Container maxWidth="sm">
      <Grid
        container
        spacing={2}
        justifyContent="center"
        alignItems="center"
        style={{ height: '100vh' }}
      >
        <Grid item xs={12}>
          <Typography variant="h4" align="center" gutterBottom>
            TODOアプリのトップページ
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <Button
            variant="contained"
            color="primary"
            component={Link}
            to="auth/login"
          >
            ログイン
          </Button>
        </Grid>
      </Grid>
    </Container>
  );
};

export default Landing;
