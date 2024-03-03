import { Edit, SimpleForm, TextInput } from 'react-admin';

const UserEdit = (props: any) => {
  return (
    <Edit {...props}>
      <SimpleForm>
        <TextInput label="ユーザー名" source="name" />
        <TextInput label="メールアドレス" source="email" type="email" />
      </SimpleForm>
    </Edit>
  );
};

export default UserEdit;
