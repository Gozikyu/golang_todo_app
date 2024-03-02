// in src/admin/index.tsx
import { Admin, Resource, ListGuesser } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import UserCreate from './UserCreate';
import { API_URL } from '@/config';

const dataProvider = jsonServerProvider(API_URL);
// const dataProvider = jsonServerProvider('https://jsonplaceholder.typicode.com');

const AdminApp = () => {
  return (
    <Admin basename="/admin" dataProvider={dataProvider}>
      <Resource name="users" list={ListGuesser} create={UserCreate} />
    </Admin>
  );
};

export default AdminApp;
