<template>
  <div class="project-view">
    <h1>Projektverwaltung</h1>
    <div class="project-form">
      <form @submit.prevent="submitProject">
        <div>
          <label for="projectName">Projektname:</label>
          <input type="text" v-model="project.name" required />
        </div>
        <div>
          <label for="projectDescription">Projektbeschreibung:</label>
          <textarea v-model="project.description" required></textarea>
        </div>
        <div>
          <label for="projectOwner">Projektinhaber:</label>
          <input type="text" v-model="project.owner" required />
        </div>
        <div>
          <label for="projectDepartment">Abteilung:</label>
          <input type="text" v-model="project.department" required />
        </div>
        <div>
          <label for="projectStatus">Status:</label>
          <select v-model="project.status">
            <option value="geplant">Geplant</option>
            <option value="aktiv">Aktiv</option>
            <option value="archiviert">Archiviert</option>
          </select>
        </div>
        <button type="submit">Projekt speichern</button>
      </form>
    </div>
    <div class="project-list">
      <h2>Projektübersicht</h2>
      <table>
        <thead>
          <tr>
            <th>Projektname</th>
            <th>Beschreibung</th>
            <th>Inhaber</th>
            <th>Abteilung</th>
            <th>Status</th>
            <th>Aktionen</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="proj in projects" :key="proj.id">
            <td>{{ proj.name }}</td>
            <td>{{ proj.description }}</td>
            <td>{{ proj.owner }}</td>
            <td>{{ proj.department }}</td>
            <td>{{ proj.status }}</td>
            <td>
              <button @click="editProject(proj.id)">Bearbeiten</button>
              <button @click="deleteProject(proj.id)">Löschen</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      project: {
        name: '',
        description: '',
        owner: '',
        department: '',
        status: 'geplant',
      },
      projects: [],
    };
  },
  methods: {
    submitProject() {
      // Logik zum Speichern des Projekts
    },
    editProject(id) {
      // Logik zum Bearbeiten des Projekts
    },
    deleteProject(id) {
      // Logik zum Löschen des Projekts
    },
    fetchProjects() {
      // Logik zum Abrufen der Projekte
    },
  },
  mounted() {
    this.fetchProjects();
  },
};
</script>

<style scoped>
.project-view {
  padding: 20px;
}
.project-form {
  margin-bottom: 20px;
}
.project-list table {
  width: 100%;
  border-collapse: collapse;
}
.project-list th, .project-list td {
  border: 1px solid #ddd;
  padding: 8px;
}
</style>