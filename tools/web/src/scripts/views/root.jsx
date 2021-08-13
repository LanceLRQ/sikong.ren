import '@/styles/index.scss';

import React from 'react';
import {
  withRouter
} from 'react-router-dom';
import { Layout } from 'antd';
import { renderRoutes } from 'react-router-config';
import {CommonHeader} from './header';

export const IndexApp = withRouter((props) => {
  const { route } = props;
  return <Layout className="site-root">
      <CommonHeader {...props} />
      <Layout className="root-layout">
        {renderRoutes(route.routes)}
      </Layout>
    </Layout>
});
