<template>
  <nav class="navbar">
    <div class="navbar-brand">
      <router-link to="/" class="navbar-item">Projektkosten Rechner</router-link>
    </div>
    <div class="navbar-menu">
      <div class="navbar-start">
        <router-link to="/projects" class="navbar-item">Projekte</router-link>
        <router-link to="/about" class="navbar-item">Über</router-link>
      </div>
      <div class="navbar-end">
        <router-link to="/login" class="navbar-item">Anmelden</router-link>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'Navigation',
}
</script>

<style scoped>
.navbar {
  background-color: #f8f9fa;
  padding: 1rem;
}

.navbar-item {
  color: #333;
  text-decoration: none;
}

.navbar-item:hover {
  text-decoration: underline;
}
</style>