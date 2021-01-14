import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SignIn from './components/SignIn'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/signin', SignIn);
  },
});
