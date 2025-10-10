<template>
  <v-sheet class="d-flex justify-content-between pa-2" v-ripple width="100%">
    <div class="text-h6 align-self-center">{{ boulderRoute.screwedDifficulty.alias }}</div>
    <v-spacer></v-spacer>
    <div class="d-flex gap-dot5rem">
      <div v-if="false" class="d-flex gap-dot5rem">
        <!--TODO: implement a collapse field, where the style is displayed-->
        <v-chip v-for="style in boulderRoute.style" :text="style.alias"></v-chip>
      </div>
      <v-icon v-if="boulderRoute.like" color="red" size="2rem">mdi-heart-outline</v-icon>
    </div>
  </v-sheet>
</template>

<script setup lang="ts">
import {Boulder} from "@/models/boulder";

defineProps<{
  boulderRoute: Boulder
}>()
</script>