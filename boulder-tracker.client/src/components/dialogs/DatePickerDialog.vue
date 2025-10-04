<template>
  <v-dialog :width="width" :height="height * 0.83">
    <template v-slot:activator="{ props: activatorProps }">
      <v-date-input v-bind="activatorProps"></v-date-input>
    </template>
    <v-date-picker :width="width" weekday-format="short" first-day-of-week="1" :header="header"
                   :max="max" :title="title" show-adjacent-months
                   v-bind:model-value="date" onclick="submit"></v-date-picker>
  </v-dialog>
</template>

<script setup lang="ts">
  defineProps<{
    width: string,
    height: number,
    title: string,
    header: string,
    max: Date,
    date: Date,
  }>()
  
  let date: Date = new Date()
  
  function submit() {
      alert(date)
  }
</script>