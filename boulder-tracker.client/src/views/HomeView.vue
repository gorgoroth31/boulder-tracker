<template>
  <div v-if="isLoading" class="w-100 h-100 d-flex justify-center align-center">
    <v-progress-linear indeterminate></v-progress-linear>
  </div>
  <div v-else>
    <div v-if="sessions === null || sessions?.length === 0">
      <v-card>
        <v-card-text @click="router.push('sessions/add')" v-ripple class="bg-primary d-flex align-center justify-content-between">
          <div>Sieht aus, als ob du noch keine Session abgespeichert hast! Klicke hier, um eine neue Session zu erstellen</div>
          <div><v-icon>mdi-chevron-right</v-icon></div>
        </v-card-text>
      </v-card>
    </div>
    <div v-else>
      <div class="text-h5">Letzte Sessions</div>
      <SessionCard v-for="session in sessions" :session="session"></SessionCard>
    </div>
  </div>
</template>

<script setup lang="ts">
import {onMounted, Ref, ref} from "vue";
import {getLatestSessions} from "@/api/session.api";
import SessionCard from "@/components/SessionCard.vue";
import router from "@/router";
import { Session } from "@/models/session";

const sessions = ref<Session[] | null>(null);

const isLoading: Ref<boolean> = ref(true);

onMounted(async () => {
  isLoading.value = true;

  await getLatestSessions().then(value => {
    sessions.value = value.data;
    isLoading.value = false;
  })
})

</script>