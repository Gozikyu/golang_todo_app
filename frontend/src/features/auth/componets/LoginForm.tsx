import { useState } from 'react';

const LoginForm = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');

  const handleLogin = async () => {
    // ユーザー認証の実際の処理を行う
    // ...

    // 認証が成功したと仮定して、ユーザー情報を保持する
    const user = { username, email };

    // ユーザー情報をローカルストレージに保存
    localStorage.setItem('user', JSON.stringify(user));
  };

  return (
    <div>
      <label>
        ユーザー名:
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
      </label>
      <br />
      <label>
        メールアドレス:
        <input
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
      </label>
      <br />
      <button onClick={handleLogin}>ログイン</button>
    </div>
  );
};

export default LoginForm;
