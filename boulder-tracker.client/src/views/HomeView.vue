<template>
  <main class="h-100 position-relative">
    <div class="dashboard-container" v-if="user !== null">
      <h1 class="text-break">Hallo {{user.userName}}</h1>
    </div>
    <AddSessionDialog btn-class="position-absolute bottom-0 w-100"></AddSessionDialog>
  </main>
</template>

<script setup lang="ts">
import AddSessionDialog from '../components/dialogs/AddSessionDialog.vue';
import {onMounted, ref} from "vue";
import {User} from "@/models/user";
import {getCurrentLoggedInUser} from "@/api/user.api";

const user = ref<User | null>(null); 

onMounted(async () => {
  await getCurrentLoggedInUser().then(value => {
    user.value = value.data;
  });
})

</script>