import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    projects: [],
    departments: [],
    user: null,
  },
  mutations: {
    SET_PROJECTS(state, projects) {
      state.projects = projects;
    },
    ADD_PROJECT(state, project) {
      state.projects.push(project);
    },
    UPDATE_PROJECT(state, updatedProject) {
      const index = state.projects.findIndex(project => project.id === updatedProject.id);
      if (index !== -1) {
        Vue.set(state.projects, index, updatedProject);
      }
    },
    DELETE_PROJECT(state, projectId) {
      state.projects = state.projects.filter(project => project.id !== projectId);
    },
    SET_DEPARTMENTS(state, departments) {
      state.departments = departments;
    },
    SET_USER(state, user) {
      state.user = user;
    },
  },
  actions: {
    fetchProjects({ commit }) {
      // API call to fetch projects
    },
    createProject({ commit }, project) {
      // API call to create a project
    },
    editProject({ commit }, project) {
      // API call to edit a project
    },
    removeProject({ commit }, projectId) {
      // API call to delete a project
    },
    fetchDepartments({ commit }) {
      // API call to fetch departments
    },
    login({ commit }, user) {
      // API call to authenticate user
    },
  },
  getters: {
    allProjects: state => state.projects,
    allDepartments: state => state.departments,
    currentUser: state => state.user,
  },
});