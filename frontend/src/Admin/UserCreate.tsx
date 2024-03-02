import * as React from 'react';
import {
  Create,
  SimpleForm,
  TextInput,
  useCreate,
  CreateProps,
} from 'react-admin';

const UserCreate: React.FC<CreateProps> = (props) => {
  const [create] = useCreate('users');

  const onSave = (values: any) => {
    create('users', { data: values });
  };

  return (
    <Create {...props}>
      <SimpleForm onSubmit={onSave}>
        <TextInput label="ユーザー名" source="name" />
        <TextInput label="メールアドレス" source="email" type="email" />
      </SimpleForm>
    </Create>
  );
};

export default UserCreate;
