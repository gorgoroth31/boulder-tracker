<template>
  <v-skeleton-loader v-if="isLoading" type="text"></v-skeleton-loader>
  <v-sheet v-else class="d-flex justify-content-between pa-2" v-ripple width="100%">
    <div class="text-body1 d-flex flex-column">
      <span>{{ startTime.toLocaleDateString() }}</span>
      <span>Dauer: {{ durationHours }}:{{durationMinutes}}</span>
    </div>
  </v-sheet>
</template>

<script setup lang="ts">

import {Session} from "@/models/session";
import {onMounted, Ref, ref} from "vue";

const isLoading: Ref<boolean> = ref(true);

const startTime: Ref<Date> = ref(new Date());
const endTime: Ref<Date> = ref(new Date());

const durationMinutes: Ref<number> = ref(0);
const durationHours: Ref<number> = ref(0);

const props = defineProps<{
  session: Session
}>()

onMounted(() => {
  startTime.value = new Date(props.session.startTime)
  endTime.value = new Date(props.session.endTime)

  const diff = endTime.value.getTime() - startTime.value.getTime();
  const totalMinutes = diff / (1000 * 60)
  
  durationHours.value = Math.floor(totalMinutes / 60)
  
  durationMinutes.value = totalMinutes % 60
  isLoading.value = false;
})

</script>