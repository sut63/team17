import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="./" text="Home" />
    <SidebarItem icon={CreateComponentIcon} to="/member" text="Member" />
    <SidebarItem icon={CreateComponentIcon} to="/student" text="Student Background" />
    <SidebarItem icon={CreateComponentIcon} to="/" text="Student Result" />
    <SidebarItem icon={CreateComponentIcon} to="/" text="Province" />
    <SidebarItem icon={CreateComponentIcon} to="/" text="Professor Background" />
    <SidebarItem icon={CreateComponentIcon} to="/" text="Student Activity" />
    <SidebarItem icon={CreateComponentIcon} to="/" text="Course" />
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarThemeToggle />
    <SidebarUserSettings />
    <SidebarPinButton />
  </Sidebar>
);