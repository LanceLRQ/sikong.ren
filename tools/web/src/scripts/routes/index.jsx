import React from 'react';
import Views from '../views';
import { CoopView } from '../views/bilibili/cooperation'

export const routes = [
  {
    component: Views.Root,
    routes: [
      {
        path: '/',
        exact: true,
        component: Views.Index,
      },
      {
        path: '/bilibili',
        component: Views.Fragment,
        routes: [
          {
            path: '/bilibili/cooperation',
            exact: true,
            component: CoopView,
          },
        ]
      },
      // {
      //   path: '/page2',
      //   component: TestClassBasedApp,
      //   fetchData: TestClassBasedApp.fetchData,
      // },
      // {
      //   path: '/typescript',
      //   component: TsApp,
      // }
    ],
  }
];
