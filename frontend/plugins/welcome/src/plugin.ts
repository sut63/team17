import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateCoursePage from './components/CreateCourse';
import SearchCoursePage from './components/SearchCourse';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/create-course', CreateCoursePage);
    router.registerRoute('/search-course', SearchCoursePage);
  },
});
