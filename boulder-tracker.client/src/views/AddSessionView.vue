<template>
  <div class="text-h5 gy-0">Wann warst du bouldern?</div>
</template>

<script setup lang="ts">
  import {onMounted} from "vue";
  import mainPageUtils from "./../utils/mainPageUtils";

  onMounted(() => {
    mainPageUtils.pageTitle.value = "Session hinzufügen";
  })
</script>