<template>
  <main class="h-100 position-relative">
    <div class="dashboard-container">
      <h1>Dashboard</h1>
      <UserProfile></UserProfile>
      <button @click="doSomethingWithToken()">stuff</button>
      <Logout></Logout>
    </div>
    <AddSessionDialog btn-class="position-absolute bottom-0 w-100"></AddSessionDialog>
  </main>
</template>

<script setup lang="ts">
import AddSessionDialog from '../components/dialogs/AddSessionDialog.vue';
import UserProfile from '../components/auth/UserProfile.vue';

import { useAuth0 } from '@auth0/auth0-vue';
import Logout from '../components/auth/Logout.vue';
import {getHealthCheck} from "@/api/api";

const { getAccessTokenSilently } = useAuth0();

async function doSomethingWithToken() {
    await getHealthCheck();
    // const token = await getAccessTokenSilently();
    // const response = await fetch('http://localhost:8080/api/health', {
    //   headers: {
    //     Authorization: 'Bearer ' + token
    //   }
    // });
    // const data = await response.json();
}

</script>