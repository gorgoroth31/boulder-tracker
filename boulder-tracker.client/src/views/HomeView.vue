<template>
      <div class="dashboard-container" v-if="user !== null">
      </div>
      <v-btn class="w-100 position-absolute bottom-0"
             color="green-darken-1"
             text="Neue Session"
             variant="flat"
             to="/sessions/add"></v-btn>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import {User} from "@/models/user";
import {getCurrentLoggedInUser} from "@/api/user.api";
import mainPageUtils from "./../utils/mainPageUtils";

const user = ref<User | null>(null); 

onMounted(async () => {
  await getCurrentLoggedInUser().then(value => {
    user.value = value.data;
    
    mainPageUtils.pageTitle.value = "Hallo " + user.value.userName
    mainPageUtils.css.value = "position-relative"
  });
})

</script>