import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import SignOut from '@material-ui/icons/Settings';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import SearchComponentIcon from '@material-ui/icons/Search';
import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';
import { Cookies } from 'plugin-welcome/src/Cookie';
import PostAddIcon from '@material-ui/icons/PostAdd';
import PersonAddIcon from '@material-ui/icons/PersonAdd';
import AddLocationIcon from '@material-ui/icons/AddLocation';
import SchoolIcon from '@material-ui/icons/School';
import GroupAddIcon from '@material-ui/icons/GroupAdd';
import BookIcon from '@material-ui/icons/Book';
var ck = new Cookies()

 function Clears() {
  ck.ClearCookie()
  window.location.reload(false)
}

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="./" text="Home" />
    <SidebarDivider />
    <SidebarItem icon={CreateComponentIcon} to="/member" text="Member" />
    <SidebarItem icon={GroupAddIcon} to="/student" text="Student Background" />
    <SidebarItem icon={BookIcon} to="/result" text="Student Result" />
    <SidebarItem icon={AddLocationIcon} to="/province" text="Province" />
    <SidebarItem icon={PersonAddIcon} to="/professor" text="Professor Background" />
    <SidebarItem icon={PostAddIcon} to="/activity" text="Student Activity" />
    <SidebarItem icon={SchoolIcon} to="/course" text="Course" />
    <SidebarDivider />
    <SidebarItem icon={SearchComponentIcon} to="/si" text="Student Info" />
    <SidebarItem icon={SearchComponentIcon} to="/sg" text="Student Grade" />
    <SidebarItem icon={SearchComponentIcon} to="/pv" text="Province Info" />
    <SidebarItem icon={SearchComponentIcon} to="/pf" text="Professor Info" />
    <SidebarItem icon={SearchComponentIcon} to="/as" text="Activity Search" />
    <SidebarItem icon={SearchComponentIcon} to="/cs" text="Course Search" />
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarThemeToggle />
    <SidebarItem
      icon={SignOut}
      to="login"
      text="Logout"
      onClick={Clears}
    />
    {/*<SidebarUserSettings />*/}
    <SidebarPinButton />
  </Sidebar>
);