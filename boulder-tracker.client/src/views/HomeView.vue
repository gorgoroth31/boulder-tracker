<template>
  <div v-if="isLoading" class="w-100 h-100 d-flex justify-center align-center">
    <v-progress-linear indeterminate ></v-progress-linear>
  </div>
  <div v-else>
    <div class="text-h5">Letzte Sessions</div>
    <SessionCard v-for="session in sessions" :session="session"></SessionCard>
  </div>
</template>

<script setup lang="ts">
import {onMounted, Ref, ref} from "vue";
import {User} from "@/models/user";
import {getCurrentLoggedInUser} from "@/api/user.api";
import mainPageUtils from "./../utils/mainPageUtils";
import {getLatestSessions} from "@/api/session.api";
import {Session} from "node:inspector";
import SessionCard from "@/components/SessionCard.vue";

const user = ref<User | null>(null);
const sessions = ref<Session[] | null>(null);

const isLoading: Ref<boolean> = ref(true);

onMounted(async () => {
  isLoading.value = true;
  await getCurrentLoggedInUser().then(value => {
    user.value = value.data;

    mainPageUtils.pageTitle.value = "Hallo " + user.value.userName
  });

  await getLatestSessions().then(value => {
    sessions.value = value.data;
  })
  isLoading.value = false;
})

</script>