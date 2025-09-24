<template>
  <Login></Login>
</template>

<script setup lang="ts">
import Login from "@/components/auth/Login.vue";
import {onMounted} from "vue";
import mainPageUtils from "@/utils/mainPageUtils";

onMounted(() => {
  mainPageUtils.pageTitle.value = "Login"
})
</script>