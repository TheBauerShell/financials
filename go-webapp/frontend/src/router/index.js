import Vue from 'vue';
import Router from 'vue-router';
import ProjectView from '../views/ProjectView.vue';

Vue.use(Router);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: ProjectView,
  },
  {
    path: '/projects',
    name: 'Projects',
    component: ProjectView,
  },
];

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;