// in src/admin/index.tsx
import { Admin, Resource, fetchUtils } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import UserCreate from './UserCreate';
import { API_URL } from '@/config';
import UserEdit from './UserEdit';
import UserList from './UserList';
import storage from '@/util/storage';

const httpClient = (url: string, options: any = {}) => {
  if (!options.headers) {
    options.headers = new Headers({ Accept: 'application/json' });
  }
  const token = storage.getToken();
  options.headers.set('Authorization', `Bearer ${token}`);
  return fetchUtils.fetchJson(url, options);
};

const dataProvider = jsonServerProvider(API_URL, httpClient);
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
