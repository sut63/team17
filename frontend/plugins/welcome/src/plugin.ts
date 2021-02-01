import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SearchGrade from './components/SearchGrade';
import Student from './components/Student'
import Logins from './components/Login'
import Member from './components/Member'
import Activity from './components/Activity'
import Result from './components/Results'
import Course from './components/Course'
import Province from './components/Province'
import Professor from './components/Professor'
import Warn from './components/Warning_Role'
import si from './components/SearchStudent'
import { Cookies } from './Cookie'

var ck = new Cookies()
var role = ck.GetRole()

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    if(role == "เจ้าหน้าที่ฝ่ายพนักงานบริการนักศึกษา"){
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/student', Student);
      router.registerRoute('/member', Member);
      router.registerRoute('/activity', Warn);
      router.registerRoute('/professor', Warn);
      router.registerRoute('/login', Warn);
      router.registerRoute('/province', Warn);
      router.registerRoute('/result', Result);
      router.registerRoute('/course', Warn);
      
      router.registerRoute('/si', si);
      router.registerRoute('/as', WelcomePage);
      router.registerRoute('/pf', WelcomePage);
      router.registerRoute('/pv', WelcomePage);
      router.registerRoute('/sg', SearchGrade);
      router.registerRoute('/cs', WelcomePage);
    }else if(role=="เจ้าหน้าที่ฝ่ายทะเบียนมหาวิทยาลัย"){
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/student', Warn);
      router.registerRoute('/member', Member);
      router.registerRoute('/activity', Warn);
      router.registerRoute('/professor', Warn);
      router.registerRoute('/login', Logins);
      router.registerRoute('/province', Warn);
      router.registerRoute('/result', Warn);
      router.registerRoute('/course', Course);

      router.registerRoute('/si', WelcomePage);
      router.registerRoute('/as', WelcomePage);
      router.registerRoute('/pf', WelcomePage);
      router.registerRoute('/pv', WelcomePage);
      router.registerRoute('/sg', WelcomePage);
      router.registerRoute('/cs', WelcomePage);
    }else if(role=="เจ้าหน้าที่ฝ่ายทะเบียน"){
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/student', Warn);
      router.registerRoute('/member', Member);
      router.registerRoute('/activity', Warn);
      router.registerRoute('/professor', Professor);
      router.registerRoute('/login', Logins);
      router.registerRoute('/province', Province);
      router.registerRoute('/result', Warn);
      router.registerRoute('/course', Warn);

      router.registerRoute('/si', WelcomePage);
      router.registerRoute('/as', WelcomePage);
      router.registerRoute('/pf', WelcomePage);
      router.registerRoute('/pv', WelcomePage);
      router.registerRoute('/sg', WelcomePage);
      router.registerRoute('/cs', WelcomePage);
    }else if(role=="เจ้าหน้าที่ฝ่ายกิจกรรมนักศึกษา"){
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/student', Warn);
      router.registerRoute('/member', Member);
      router.registerRoute('/activity', Activity);
      router.registerRoute('/professor', Professor);
      router.registerRoute('/login', Logins);
      router.registerRoute('/province', Warn);
      router.registerRoute('/result', Warn);
      router.registerRoute('/course', Warn);

      router.registerRoute('/si', WelcomePage);
      router.registerRoute('/as', WelcomePage);
      router.registerRoute('/pf', WelcomePage);
      router.registerRoute('/pv', WelcomePage);
      router.registerRoute('/sg', WelcomePage);
      router.registerRoute('/cs', WelcomePage);
    }else{
      router.registerRoute('/', Logins);
      router.registerRoute('/student', Logins);
      router.registerRoute('/member', Logins);
      router.registerRoute('/activity', Logins);
      router.registerRoute('/professor', Logins);
      router.registerRoute('/login', Logins);
      router.registerRoute('/province', Logins);
      router.registerRoute('/result', Logins);
      router.registerRoute('/course', Logins);

      router.registerRoute('/si', Logins);
      router.registerRoute('/as', Logins);
      router.registerRoute('/pf', Logins);
      router.registerRoute('/pv', Logins);
      router.registerRoute('/sg', Logins);
      router.registerRoute('/cs', Logins);
    }
  },
});
