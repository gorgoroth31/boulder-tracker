<template>
  <main class="h-100 position-relative">
    <div class="dashboard-container">
      <h1>Dashboard</h1>
      <UserProfile></UserProfile>
    </div>
    <AddSessionDialog btn-class="position-absolute bottom-0 w-100"></AddSessionDialog>
  </main>
</template>

<script setup lang="ts">
import AddSessionDialog from '../components/dialogs/AddSessionDialog.vue';
import UserProfile from "@/components/auth/UserProfile.vue";
</script>