<template>
  <v-carousel
      height="300"
      show-arrows="hover"
      cycle
      hide-delimiter-background>
    <v-carousel-item
        v-for="(slide, i) in slides"
        :key="i">
      <v-sheet
          color="dark-grey"
          height="100%">
        <div class="d-flex fill-height justify-center align-center">
          <div class="text-h4 text-center">
            {{ slide }}
          </div>
        </div>
      </v-sheet>
    </v-carousel-item>
  </v-carousel>
</template>

<script setup lang="ts">
defineProps<{
  slides: string[]
}>()
</script>