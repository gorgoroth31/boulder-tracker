<script setup lang="ts">
import { RouterView} from 'vue-router'
import mainPageUtil from "./utils/mainPageUtils";
</script>

<template>
  <main :class="'h-100 overflow-scroll flex flex-col md:flex-row gap-4 p-6 ' + mainPageUtil.css.value">
    <div class="text-h4 text-break">{{ mainPageUtil.pageTitle }}</div>
    <RouterView/>
  </main>
</template>

<style scoped>

</style>
