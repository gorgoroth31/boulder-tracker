<template>
  <div class="text-h5">Hier kommt eine putzig kleine About-page hin</div>
</template>

<script setup lang="ts">
  import {onMounted} from "vue";
  import mainPageUtils from "@/utils/mainPageUtils";

  onMounted(() => {
    mainPageUtils.pageTitle.value = "Über"
  })
</script>

<style scoped>

</style>