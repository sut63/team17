import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import SearchOutlined from '@material-ui/icons/SearchOutlined';
import BookOutlined from '@material-ui/icons/BookOutlined';
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
    <SidebarItem icon={BookOutlined} to="./create-course" text="บันทึกข้อมูลหลักสูตร" />
    <SidebarItem icon={SearchOutlined} to="./search-course" text="ข้อมูลหลักสูตร" />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" />
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarThemeToggle />
    <SidebarUserSettings />
    <SidebarPinButton />
  </Sidebar>
);