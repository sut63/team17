import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Student from './components/Student'
import Logins from './components/Login'
import Member from './components/Member'
import Activity from './components/Activity'
import Result from './components/Results'
import Course from './components/Course'
import Province from './components/Province'
import Professor from './components/Professor'
import { Cookies } from './Cookie'

var ck = new Cookies()
var role = ck.GetRole()

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if(role != "ทะเบียน"){
      router.registerRoute('/', Logins);
      router.registerRoute('/student', Logins);
      router.registerRoute('/member', Logins);
      router.registerRoute('/activity', Logins);
      router.registerRoute('/professor', Logins);
      router.registerRoute('/login', Logins);
      router.registerRoute('/province', Logins);
      router.registerRoute('/result', Logins);
      router.registerRoute('/course', Logins);
    }else{
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/student', Student);
      router.registerRoute('/member', Member);
      router.registerRoute('/activity', Activity);
      router.registerRoute('/professor', Professor);
      router.registerRoute('/login', Logins);
      router.registerRoute('/province', Province);
      router.registerRoute('/result', Result);
      router.registerRoute('/course', Course);
    }
  },
});
