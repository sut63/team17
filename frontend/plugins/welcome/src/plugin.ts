import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Logins from './components/Login';
import Member from './components/Member';
import Warn from './components/Warning_Role';

import Student from './components/Student';
import Activity from './components/Activity';
import Result from './components/Results';
import Course from './components/Course';
import Province from './components/Province';
import Professor from './components/Professor';

import SearchStudent from './components/SearchStudent';
import SearchGrade from './components/SearchGrade';
import SearchActivity from './components/SearchActivity';
import SearchCourse from './components/SearchCourse';
import SearchProfessor from './components/SearchProfessor';
import SearchProvince from './components/SearchProvince';

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
      
      router.registerRoute('/si', SearchStudent);
      router.registerRoute('/as', Warn);
      router.registerRoute('/pf', SearchProfessor);
      router.registerRoute('/pv', Warn);
      router.registerRoute('/sg', SearchGrade);
      router.registerRoute('/cs', Warn);
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

      router.registerRoute('/si', Warn);
      router.registerRoute('/as', Warn);
      router.registerRoute('/pf', SearchProfessor);
      router.registerRoute('/pv', Warn);
      router.registerRoute('/sg', Warn);
      router.registerRoute('/cs', SearchCourse);
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
      router.registerRoute('/pf', SearchProfessor);
      router.registerRoute('/si', Warn);
      router.registerRoute('/as', Warn);
     
      router.registerRoute('/pv', SearchProvince);
      router.registerRoute('/sg', Warn);
      router.registerRoute('/cs', Warn);
    }else if(role=="เจ้าหน้าที่ฝ่ายกิจกรรมนักศึกษา"){
      router.registerRoute('/', WelcomePage);
      router.registerRoute('/student', Warn);
      router.registerRoute('/member', Member);
      router.registerRoute('/activity', Activity);
      router.registerRoute('/professor', Warn);
      router.registerRoute('/login', Logins);
      router.registerRoute('/province', Warn);
      router.registerRoute('/result', Warn);
      router.registerRoute('/course', Warn);

      router.registerRoute('/si', Warn);
      router.registerRoute('/as', SearchActivity);
      router.registerRoute('/pf', SearchProfessor);
      router.registerRoute('/pv', Warn);
      router.registerRoute('/sg', Warn);
      router.registerRoute('/cs', Warn);
      
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
      //router.registerRoute('/pf', Logins);
      router.registerRoute('/pv', Logins);
      router.registerRoute('/sg', Logins);
      router.registerRoute('/cs', Logins);
      router.registerRoute('/pf', SearchProfessor);
    }
  },
});
