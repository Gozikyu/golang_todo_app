// in src/admin/index.tsx
import { Admin, Resource } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import UserCreate from './UserCreate';
import { API_URL } from '@/config';
import UserEdit from './UserEdit';
import UserList from './UserList';

const dataProvider = jsonServerProvider(API_URL);
// const dataProvider = jsonServerProvider('https://jsonplaceholder.typicode.com');

const AdminApp = () => {
  return (
    <Admin basename="/admin" dataProvider={dataProvider}>
      <Resource
        name="users"
        list={UserList}
        edit={UserEdit}
        create={UserCreate}
      />
    </Admin>
  );
};

export default AdminApp;
