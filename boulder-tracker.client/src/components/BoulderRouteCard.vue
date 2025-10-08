<template>
  <v-sheet class="d-flex justify-content-between pa-2" v-ripple width="100%">
    {{ boulderRoute.screwedDifficulty.alias }}
    <v-spacer></v-spacer>
    <v-icon color="red" size="2rem">mdi-heart-outline</v-icon>
  </v-sheet>
</template>

<script setup lang="ts">
import {Boulder} from "@/models/boulder";

defineProps<{
  boulderRoute: Boulder
}>()
</script>

<style scoped>

</style>