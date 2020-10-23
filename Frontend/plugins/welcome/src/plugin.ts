import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/Users';
import Table from './components/Table';
import Login from './components/Login';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/Create', CreateUser);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/table', Table);
    
  }
});
