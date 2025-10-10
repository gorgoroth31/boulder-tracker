<template>
  <v-expansion-panel>
    <v-expansion-panel-title class="pa-0">
      <v-sheet class="d-flex justify-content-between pa-2" v-ripple width="100%">
        <div class="text-h6 align-self-center">{{ boulderRoute.screwedDifficulty.alias }}</div>
        <v-spacer></v-spacer>
      </v-sheet>
    </v-expansion-panel-title>
    <v-expansion-panel-text class="no-inner-padding">
      <div class="px-3 pb-3">
        <div class="d-flex flex-wrap gap-dot5rem mb-3">
          <v-chip v-for="style in boulderRoute.style" :text="style.alias"></v-chip>
        </div>
        <div class="d-flex">
          <v-icon :color="boulderRoute.like ? 'red' : 'surfaceLighter'" size="2rem">mdi-heart-outline</v-icon>
          <v-spacer></v-spacer>
          <v-icon v-if="!areDeleteButtonsOpen" @click="areDeleteButtonsOpen = true" size="2rem">mdi-delete-outline</v-icon>
          <div class="d-flex gap-1rem" v-else>
            <v-icon @click="handleDelete" size="2rem">mdi-check</v-icon>
            <v-icon @click="areDeleteButtonsOpen = false" size="2rem">mdi-close</v-icon>
          </div>
        </div>
      </div>
    </v-expansion-panel-text>
  </v-expansion-panel>
</template>

<script setup lang="ts">
import {Boulder} from "@/models/boulder";
import {ref, Ref} from "vue";

const areDeleteButtonsOpen : Ref<boolean> = ref(false);

const emit = defineEmits<{
  (e: 'delete', route: Boulder): void,
}>()

const props = defineProps<{
  boulderRoute: Boulder
}>()

function handleDelete () {
  emit('delete', props.boulderRoute)
}

</script>