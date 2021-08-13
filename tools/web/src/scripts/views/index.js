export default {
  Index: require('./welcome').Welcome,
  Root: require('./root').IndexApp,
  Bilibili: require('./bilibili/index').default,
  Fragment: require('./fragment').RouterFragment
}
