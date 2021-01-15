import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Student from './components/Student'
import Logins from './components/Login'
import Member from './components/Member'
import Results from './components/Results'
import { Cookies } from './Cookie'
import Activity from './components/Activity';

var ck = new Cookies()
var role = ck.GetRole()

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if(role != "ทะเบียน"){
      router.registerRoute('/', Logins);
    }else{
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/student', Student);
      router.registerRoute('/member', Member);
      router.registerRoute('/result', Results);
      router.registerRoute('/activity', Activity);
    }
  },
});
