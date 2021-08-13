import { Layout, Menu } from 'antd';
import React from 'react';
import { matchPath } from 'react-router';
import { RouteUrls } from '../routes/urls';

const { Header } = Layout;

const MenuItems = [
  {
    title: '首页',
    routes: RouteUrls.Index,  // like /users/:id; either a single string or an array of strings
    exact: true,
    link: RouteUrls.Index,
    key: 'index',
  },
  {
    title: 'B站工具',
    routes: RouteUrls.Bilibili.Index,
    link: RouteUrls.Bilibili.Index,
    key: 'bilibili',
  },
];


export const CommonHeader =({ location, history }) => {
  const currentUrl = location.pathname;

  const handleNavItemClick = (path, key) => (e) => {
    if (e) {
      e.preventDefault && e.preventDefault();
      e.stopPropagation && e.stopPropagation();
      e.returnValue = false;
    }
    if (key === 'register' || key === 'login') {
      // const currentUrl = getCurrentURL();
      // const query = getQuery(location);
      // if (!query.from) {
      //   history.push(buildPath(path, null, null, { from: currentUrl }));
      // } else {
      //   history.push(buildPath(path, null, null, query));
      // }
    } else {
      history.push(path);
    }
    return false;
  };

  const matchedItem = MenuItems.filter(item => !!item)
    .filter((item) => {
      return !!matchPath(currentUrl, {
        path: item.routes,
        exact: item.exact || false,
        strict: item.strict ||  false,
      });
    });

  return <Header className="header">
    <span className="logo">司空人</span>
    <Menu theme="dark" mode="horizontal" selectedKeys={matchedItem.map(item => item.key)}>
      {MenuItems.map(item => {

        return <Menu.Item key={item.key} onClick={handleNavItemClick(item.link, item.key)}>{item.title}</Menu.Item>
      })}
    </Menu>
  </Header>
}
